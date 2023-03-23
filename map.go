package collection

import "sync"

type KV[K comparable, T any] struct {
	Key   K
	Value T
}

func MapKeys[K comparable, T any](source map[K]T) []K {
	var result = make([]K, 0, len(source))
	for key := range source {
		result = append(result, key)
	}

	return result
}

func MapValues[K comparable, T any](source map[K]T) []T {
	var result = make([]T, 0, len(source))
	for _, value := range source {
		result = append(result, value)
	}

	return result
}

type SyncMap[K comparable, V any] struct {
	m sync.Map
}

func (m *SyncMap[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	if v, ok := m.m.Load(key); ok {
		return v.(V), ok
	}

	return value, false
}

func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	a, loaded := m.m.LoadOrStore(key, value)
	return a.(V), loaded
}

func (m *SyncMap[K, V]) Delete(key K) {
	m.m.Delete(key)
}

func (m *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	if v, loaded := m.m.LoadAndDelete(key); loaded {
		return v.(V), loaded
	}

	return value, false
}

func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func (m *SyncMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	if v, loaded := m.m.Swap(key, value); loaded {
		return v.(V), loaded
	}

	return previous, false
}

func (m *SyncMap[K, V]) CompareAndSwap(key K, old V, new V) bool {
	return m.m.CompareAndSwap(key, old, new)
}

func (m *SyncMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return m.m.CompareAndDelete(key, old)
}
