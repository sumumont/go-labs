package utils

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)
import "sync"

func TestKeyMutex(t *testing.T) {
	lock := New()
	lock.Lock("hysen")
	lock.UnLock("hysen")
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			lock.Lock("hysen")
			time.Sleep(time.Second * time.Duration(1))
			fmt.Println("lock", i)
			defer func() {
				lock.UnLock("hysen")
				fmt.Println("unlock", i)
			}()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("end")
}
func test(err error) (string, error) {
	//return errors.New("a1 error")
	if err == nil {
		return "nil", nil
	}
	return "", err
}

type ListModel struct {
	sync.Mutex
	Str   []int
	Error error
}

func TestMutex(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(10)
	model := ListModel{
		Mutex: sync.Mutex{},
		Str:   []int{},
		Error: nil,
	}
	for i := 0; i < 10; i++ {
		go func(i int, model *ListModel) {
			defer wg.Done()
			var err error
			defer func() {
				model.Lock()
				model.Str = append(model.Str, i)
				if err != nil {
					model.Error = err
				}
				model.Unlock()
			}()
			//if i == 3 {
			//	_, err = test(errors.New("new error 3"))
			//}
			if err != nil {
				fmt.Println(err.Error())
			}
		}(i, &model)
	}
	fmt.Println("wait")
	wg.Wait()
	fmt.Println(model)
	if model.Error != nil {
		fmt.Println(model.Error.Error())
	}
}

func TestErrorMutiFuc(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(10)
	model := ListModel{}
	var err error
	ctx, cancel := context.WithCancel(context.Background())
	go func(model *ListModel, err error, ctx context.Context, cancel context.CancelFunc) {
		for i := 0; i < 10; i++ {
			go func(i int, model *ListModel, err error, ctx context.Context, cancel context.CancelFunc) {
				defer func() {
					time.Sleep(2 * time.Second)
					fmt.Println("wg.done")
					wg.Done()
				}()
				err = a()
				if err != nil {
					fmt.Println(err.Error())
					cancel()
					return
				}
				err = b()
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				err = c()
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				model.Lock()
				model.Str = append(model.Str, i)
				model.Unlock()
			}(i, model, err, ctx, cancel)
		}
		select {
		case <-ctx.Done():
			fmt.Println("ctx done")
		}
	}(&model, err, ctx, cancel)
	wg.Wait()
	if err != nil {
		fmt.Println("result err:", err.Error())
	}
	fmt.Println(model)
}

type Result struct {
	Error error
	Count int
}

func GetResult(done chan struct{}, results chan Result, urls ...string) {
	fmt.Println("make(chan Result)")
	go func() {
		for i, url := range urls {
			go func(url string, i int) {
				fmt.Println("url:", url)
				time.Sleep(time.Second * 2)
				var result Result
				err := a()
				if err != nil {
					return
				}
				result = Result{
					Error: err,
					Count: 0,
				}
				results <- result
			}(url, i)
		}
		select {
		case <-done:
			fmt.Println("return1")
			return
		}
	}()
}

func TestGetResult(t *testing.T) {
	done := make(chan struct{})
	urls := []string{"a", "https://www.baidu.com", "b", "c", "d"}
	results := make(chan Result)
	GetResult(done, results, urls...)
	for i := 0; i < len(urls); i++ {
		select {
		case result, ok := <-results:
			if ok {
				fmt.Println("error:", result.Error, "counts:", result.Count)
			} else {
				fmt.Println("chan closed")
			}
		}
	}
	//for result := range results{
	//	fmt.Println("error:",result.Error,"counts:",result.Count)
	//}
}

func a() error {
	//return errors.New("a1 error")
	return nil
}

func b() error {
	return errors.New("b1 error")
}
func c() error {
	return errors.New("c1 error")
}
