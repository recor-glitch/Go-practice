package singleSafeMap

import (
	"sync"
)

type safeMap struct {
	mux sync.RWMutex
	m   map[string]any
}

func NewSafeMap() *safeMap {
	return &safeMap{
		m: make(map[string]any),
	}
}

func (sm *safeMap) Get(key string) any {
	sm.mux.RLock()
	defer sm.mux.RUnlock()

	value := sm.m[key]
	return value
}

func (sm *safeMap) Set(key string, value any) {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	sm.m[key] = value
}

func (sm *safeMap) Delete(key string) {
	sm.mux.Lock()
	defer sm.mux.Unlock()

	delete(sm.m, key)
}
