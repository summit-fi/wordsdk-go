package word

import "sync"

func NewMap[K comparable, V any]() Map[K, V] {
	return Map[K, V]{
		m:  make(map[K]V),
		mu: sync.RWMutex{},
	}
}

func (m *Map[K, V]) IsInitialized() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m != nil
}

type Map[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

func (m *Map[K, V]) Get(key K) V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m[key]
}

func (m *Map[K, V]) GetAll() []K {
	m.mu.RLock()
	defer m.mu.RUnlock()
	keys := make([]K, 0, len(m.m))
	for k := range m.m {
		keys = append(keys, k)
	}
	return keys
}

func (m *Map[K, V]) Set(key K, value V) {
	m.mu.Lock()
	m.m[key] = value
	m.mu.Unlock()
}

func (m *Map[K, V]) Delete(key K) {
	m.mu.Lock()
	delete(m.m, key)
	m.mu.Unlock()
}

func (m *Map[K, V]) Exist(key K) (value V, ok bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, ok = m.m[key]
	return
}
