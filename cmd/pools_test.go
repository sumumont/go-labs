package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-labs/internal/logging"
	"github.com/panjf2000/ants/v2"
	"path"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type TaskItem struct {
	Wg     *sync.WaitGroup
	Result *int32
	I      int32
	Done   *chan struct{}
	Ctx    context.Context
}

func myFunc(i interface{}) (err error) {
	task := i.(*TaskItem)
	defer func() {
		task.Wg.Done()
	}()
	atomic.AddInt32(task.Result, task.I)
	fmt.Printf("run with %d\n", task.I)
	err = errors.New("failed")
	return
}

func run(p *ants.PoolWithFunc, wg *sync.WaitGroup, runTimes int, result *int32) {
	for i := 1; i <= runTimes; i++ {
		wg.Add(1)
		task := &TaskItem{
			Wg:     wg,
			Result: result,
			I:      int32(i),
		}
		err := p.Invoke(task)
		if err != nil {
			panic(err)
		}
	}
	wg.Wait()
	//time.Sleep(5 * time.Second)
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", *result)
}

func TestRun(t *testing.T) {
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
	}, ants.WithExpiryDuration(3*time.Second))
	defer p.Release()
	var wg sync.WaitGroup
	var result int32
	run(p, &wg, 100, &result)
}

func TestTimeParse(t *testing.T) {
	fmt.Println(isValidImageName("20190610173026_1560159026665.jpg"))
}
func isValidImageName(filePath string) bool {
	filename := path.Base(filePath)
	suffix := path.Ext(filename)
	nameOnly := filePath[:len(filePath)-len(suffix)]

	splits := strings.Split(nameOnly, "_")
	if len(splits) != 2 {
		return false
	}

	res, err := strconv.ParseInt(splits[1], 10, 64)
	if err != nil {
		logging.Warn().Msgf("un support image name %s", err.Error())
		return false
	}
	t := time.UnixMilli(res)

	format := t.Format("20060102150405")

	return format == splits[0]

	//t, err := time.ParseInLocation("20060102150405", splits[0], time.Local)
	//if err != nil {
	//	logging.Warn().Msgf("un support image name %s", err.Error())
	//	return false
	//}

	//return t.UnixMilli() == (res / 1000 * 1000)
}
