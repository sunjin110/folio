package smap

import "sync"

type Map[K comparable, V any] struct {
	mutex sync.RWMutex
	m     map[K]V
}

func (m *Map[K, V]) Put(k K, v V) {
	m.mutex.Lock()
	m.m[k] = v
	m.mutex.Unlock()
}

func (m *Map[K, V]) Get(k K) (V, bool) {
	m.mutex.RLock()
	v, ok := m.m[k]
	m.mutex.RUnlock()
	return v, ok
}

func (m *Map[K, V]) GetRawMap() map[K]V {
	return m.m
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		mutex: sync.RWMutex{},
		m:     make(map[K]V),
	}
}
