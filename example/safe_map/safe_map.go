package safe_map

import (
	"errors"
	"fmt"
	"sync"
)

type SafeMap[K comparable, V any] struct {
	m    map[K]V
	lock sync.RWMutex
	zero V
}

func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{m: make(map[K]V)}
}

func (s *SafeMap[K, V]) LoadOrStore(key K, val V) (V, bool) {
	oldVal, ok := s.Get(key)
	if ok {
		return oldVal, true
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	oldVal, ok = s.m[key] // double-check
	if ok {
		return oldVal, true
	}
	s.m[key] = val
	return val, false
}

func (s *SafeMap[K, V]) Get(key K) (V, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	val, ok := s.m[key]
	return val, ok
}

func (s *SafeMap[K, V]) Put(key K, newVal V) (V, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := s.m[key]
	if !ok {
		exc := fmt.Sprintf("the key %v not exists", key)
		return s.zero, errors.New(exc)
	}
	s.m[key] = newVal
	return newVal, nil
}

func (s *SafeMap[K, V]) Delete(key K) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := s.m[key]
	if !ok {
		exc := fmt.Sprintf("the key %v not exists", key)
		return errors.New(exc)
	}
	delete(s.m, key)
	return nil
}
