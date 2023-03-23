package collection

// Contains returns true if the given item is present in the slice.
func Contains[T comparable](source []T, item T) bool {
	for _, v := range source {
		if v == item {
			return true
		}
	}

	return false
}

// MapContains returns true if the given key is present in the map.
func MapContains[K comparable, T any](source map[K]T, item K) bool {
	_, ok := source[item]
	return ok
}

// Any: Returns true if at least one element in the slice satisfies the given predicate function.
func Any[T any](source []T, predicate func(T) bool) bool {
	for _, v := range source {
		if predicate(v) {
			return true
		}
	}

	return false
}
