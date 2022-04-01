package collection

func GroupBy[T any, K comparable](source []T, keyFunc func(T) K) map[K][]T {
	var result = make(map[K][]T)

	for _, v := range source {
		var key = keyFunc(v)

		result[key] = append(result[key], v)
	}

	return result
}
