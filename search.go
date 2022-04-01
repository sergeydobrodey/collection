package collection

func Contains[T comparable](source []T, item T) bool {
	for _, v := range source {
		if v == item {
			return true
		}
	}

	return false
}

func MapContains[K comparable, T any](source map[K]T, item K) bool {
	_, ok := source[item]
	return ok
}

func Any[T any](source []T, predicate func(T) bool) bool {
	for _, v := range source {
		if predicate(v) {
			return true
		}
	}

	return false
}
