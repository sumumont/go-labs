package utils

import (
	"sync"
)

type KeyMutex struct {
	mutexes *sync.Map // Zero value is empty and ready for use
}

func (m *KeyMutex) Lock(key string) {
	value, _ := m.mutexes.LoadOrStore(key, &sync.Mutex{})
	mtx := value.(*sync.Mutex)
	mtx.Lock()
}

func (m *KeyMutex) UnLock(key string) {
	value, _ := m.mutexes.Load(key)
	if value == nil {
		return
	}
	mtx := value.(*sync.Mutex)
	mtx.Unlock()
}

func New() *KeyMutex {
	return &KeyMutex{mutexes: &sync.Map{}}
}
