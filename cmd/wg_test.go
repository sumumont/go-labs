package main

import (
	"github.com/go-labs/internal/logging"
	"sync"
	"testing"
	"time"
)

type Stuff struct {
	Id             int  `json:"id"`
	HasInferResult bool `json:"hasInferResult"`
}

func TestWg(t *testing.T) {
	list := []*Stuff{}
	wg := sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		pbS := &Stuff{
			Id:             i,
			HasInferResult: false,
		}
		wg.Add(1)
		go stuffHasInfer(&wg, pbS)
		list = append(list, pbS)
	}
	logging.Debug().Interface("list", list).Send()
	wg.Wait()
	logging.Debug().Interface("list", list).Send()
}
func stuffHasInfer(wg *sync.WaitGroup, p *Stuff) {
	defer wg.Done()
	time.Sleep(time.Second * 3)
	if p.Id%2 == 0 {
		p.HasInferResult = true
	}
}
