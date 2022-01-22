package utils

import (
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
