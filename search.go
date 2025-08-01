package collection

// Contains returns true if the given item is present in the slice.
func Contains[S ~[]T, T comparable](source S, item T) bool {
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
func Any[S ~[]T, T any](source S, predicate func(T) bool) bool {
	for _, v := range source {
		if predicate(v) {
			return true
		}
	}

	return false
}

// All: Returns true if every element in the slice satisfies the given predicate function.
func All[S ~[]T, T any](source S, predicate func(T) bool) bool {
	for _, v := range source {
		if !predicate(v) {
			return false
		}
	}

	return true
}
