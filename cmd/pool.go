package main

import (
	"fmt"
	"github.com/go-labs/internal/logging"
	"sync"
	"time"
)

type PoolTask interface {
	Do() error
}
type PoolAnalysis struct {
	Total   int `json:"total"`
	Success int `json:"success"`
	Failed  int `json:"failed"`
}
type Pool struct {
	error    *MultiError
	analysis PoolAnalysis
}

func (pool *Pool) Run(tasks []PoolTask) {
	defer TimeCost()()
	if pool.error == nil {
		pool.error = &MultiError{
			Code:   -1,
			Errors: nil,
			Mutex:  sync.Mutex{},
		}
	}
	errs := &MultiError{
		Code:   -1,
		Errors: nil,
		Mutex:  sync.Mutex{},
	}
	waiter := sync.WaitGroup{}
	waiter.Add(len(tasks))
	for i, _ := range tasks {
		go func(i int) {
			fmt.Println("i: ", i)
			defer func() {
				waiter.Done()
			}()
			err := tasks[i].Do()
			if err != nil {
				errs.Lock()
				errs.Append(err)
				errs.Unlock()
			}
		}(i)
	}
	waiter.Wait()
	pool.error.Append(errs.Errors...)

	pool.analysis.Failed = pool.analysis.Failed + len(errs.Errors)
	pool.analysis.Total = pool.analysis.Total + len(tasks)
	pool.analysis.Success = pool.analysis.Success + len(tasks) - len(errs.Errors)
}

func (pool *Pool) Error() (*PoolAnalysis, *MultiError) {
	if pool.error == nil {
		return nil, nil
	}
	if pool.error.IsErr() {
		return &pool.analysis, pool.error
	}
	return &pool.analysis, nil
}
func TimeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		logging.Debug().Msgf("time const:%v", tc)
	}
}
