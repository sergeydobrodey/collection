package collection

type KV[K comparable, T any] struct {
	Key K
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
