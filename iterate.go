package collection

// Each calls the given function for each element in the slice.
func Each[S ~[]T, T any](source S, do func(T)) {
	for _, v := range source {
		do(v)
	}
}

// MapEach calls the given function for each key-value pair in the map.
func MapEach[K comparable, T any](source map[K]T, do func(key K, value T)) {
	for k, v := range source {
		do(k, v)
	}
}
