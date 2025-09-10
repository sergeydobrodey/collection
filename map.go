package collection

import (
	"sync"
)

type Pair[T any, V any] struct {
	First  T
	Second V
}

// KV represents a key-value pair from a map operation.
type KV[K comparable, T any] struct {
	Key   K
	Value T
}

// MapKeys returns a new slice containing all keys in the map.
func MapKeys[K comparable, T any](source map[K]T) []K {
	var result = make([]K, 0, len(source))
	for key := range source {
		result = append(result, key)
	}

	return result
}

// MapValues returns a slice of type T containing the values of the source map of type T. Order is not guaranteed.
func MapValues[K comparable, T any](source map[K]T) []T {
	var result = make([]T, 0, len(source))
	for _, value := range source {
		result = append(result, value)
	}

	return result
}

// SyncMap is a thread-safe map with type-safe operations using generics.
// It wraps sync.Map and provides type safety for keys and values.
type SyncMap[K comparable, V any] struct {
	m sync.Map
}

// Store sets the value for a key.
func (m *SyncMap[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

// Load returns the value stored in the map for a key, or zero value if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	if v, ok := m.m.Load(key); ok {
		return v.(V), ok
	}

	return value, false
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	a, loaded := m.m.LoadOrStore(key, value)
	return a.(V), loaded
}

// Delete deletes the value for a key.
func (m *SyncMap[K, V]) Delete(key K) {
	m.m.Delete(key)
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	if v, loaded := m.m.LoadAndDelete(key); loaded {
		return v.(V), loaded
	}

	return value, false
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration. Read sync.Map Range for more details
func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

// Swap swaps the value for a key and returns the previous value if any.
// The loaded result reports whether the key was present.
func (m *SyncMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	if v, loaded := m.m.Swap(key, value); loaded {
		return v.(V), loaded
	}

	return previous, false
}

// CompareAndSwap swaps the old and new values for key
// if the value stored in the map is equal to old.
// The old value must be of a comparable type.
func (m *SyncMap[K, V]) CompareAndSwap(key K, old V, new V) bool {
	return m.m.CompareAndSwap(key, old, new)
}

// CompareAndDelete deletes the entry for key if its value is equal to old.
// The old value must be of a comparable type.
func (m *SyncMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return m.m.CompareAndDelete(key, old)
}

// MapFirst returns the first key-value pair from the map that satisfies the given predicate function.
// Since map iteration order is not guaranteed, "first" means any matching element.
// The ok result indicates whether a matching element was found in the map.
func MapFirst[K comparable, T any](source map[K]T, predicate func(key K, value T) bool) (result KV[K, T], ok bool) {
	for key, value := range source {
		if predicate(key, value) {
			return KV[K, T]{Key: key, Value: value}, true
		}
	}

	return result, false
}

