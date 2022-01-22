package utils

import (
	"fmt"
	"sync"
)

type KeyMutex struct {
	mutexes *sync.Map // Zero value is empty and ready for use
}

func (m *KeyMutex) Lock(key string) {
	value, _ := m.mutexes.LoadOrStore(key, &sync.Mutex{})
	mtx := value.(*sync.Mutex)
	mtx.Lock()
	fmt.Println("lock:", key)
}

func (m *KeyMutex) UnLock(key string) {
	value, _ := m.mutexes.Load(key)
	mtx := value.(*sync.Mutex)
	mtx.Unlock()
	fmt.Println("unlock:", key)
}

func New() *KeyMutex {
	return &KeyMutex{mutexes: &sync.Map{}}
}
