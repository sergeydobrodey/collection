package collection

// GroupBy groups the elements of the slice by a key returned by the given key function.
func GroupBy[S ~[]T, T any, K comparable](source S, keyFunc func(T) K) map[K]S {
	var result = make(map[K]S)

	for _, v := range source {
		var key = keyFunc(v)

		result[key] = append(result[key], v)
	}

	return result
}
