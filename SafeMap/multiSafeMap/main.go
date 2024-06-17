package multiSafeMap

import (
	"fmt"
	"sync"

	"github.com/recor-glitch/Go-practice/hash"
)

type SafeMap struct {
	mux sync.RWMutex
	m   map[string]any
}

type multiSafeMap []*SafeMap

func NewMultiSafeMap(cap int) *multiSafeMap {
	myMultiSafeMap := make(multiSafeMap, cap)
	for i := range myMultiSafeMap {
		myMultiSafeMap[i] = &SafeMap{
			m: make(map[string]any),
		}
	}
	return &myMultiSafeMap
}

func (m multiSafeMap) Set(key string, value any) (success bool, err error) {
	hashedKey, hashErr := hash.GetHash(key)
	if hashErr != nil {
		return false, hashErr
	}
	index, er := hash.GetMapIndexFromHash(hashedKey, len(m))
	if er != nil {
		return false, er
	}
	myMap := m[index]
	myMap.mux.Lock()
	defer myMap.mux.Unlock()
	myMap.m[key] = value
	return true, nil
}

func (m multiSafeMap) Get(key string) (any, error) {
	haskedKey, hashErr := hash.GetHash(key)
	if hashErr != nil {
		return nil, hashErr
	}
	hashedIndex, err := hash.GetMapIndexFromHash(haskedKey, len(m))
	if err != nil {
		return nil, err
	}
	value, ok := m[hashedIndex].m[key]
	if !ok {
		return nil, fmt.Errorf("value is nil")
	}
	fmt.Printf("Get func value: %v\n", value)
	fmt.Printf("Get func array: %v\n", m)
	return value, nil
}

func (m multiSafeMap) Delete(key string) (bool, error) {
	haskedKey, hashErr := hash.GetHash(key)
	if hashErr != nil {
		return false, hashErr
	}
	hashedIndex, err := hash.GetMapIndexFromHash(haskedKey, len(m))
	if err != nil {
		return false, err
	}
	_, ok := m[hashedIndex].m[key]
	if !ok {
		return false, fmt.Errorf("key %v not found", key)
	}
	delete(m[hashedIndex].m, key)
	return true, nil
}

func (m multiSafeMap) Display() {
	for _, mm := range m {
		fmt.Printf("%+v", mm)
	}
}
