package main

import (
	"fmt"
	"testing"
	"time"
)

func TestPoolTask(t *testing.T) {
	var list []Job
	for i := 0; i < 5; i++ {
		list = append(list, Job{Num: i})
	}

	itoa, err := IteratorNew(2, list)
	if err != nil {
		panic(err)
	}
	pool := Pool{}
	for {
		newlist := itoa.Next()
		if newlist == nil {
			break
		}
		jobss := newlist.([]Job)
		fmt.Println(jobss)
		var jobs []PoolTask
		for i, _ := range jobss {
			job := &jobss[i]
			jobs = append(jobs, job)
		}
		pool.Run(jobs)
	}
	result, err := pool.Error()
	fmt.Println(result)
	fmt.Println(err)
}

type Job struct {
	Num int
}

func (task Job) Do() error {
	//if task.Num%2 == 1 {
	//	return errors.New("can not been eve")
	//}
	fmt.Println(task.Num)
	time.Sleep(1 * time.Second)
	return nil
}
