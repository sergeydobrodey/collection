package collection

func TransformBy[T, K any](source []T, transform func(T) K) []K {
	var result = make([]K, len(source))
	for i, item := range source {
		result[i] = transform(item)
	}

	return result
}

func TransformManyBy[T, K any](source []T, transform func(T) []K) []K {
	var many = TransformBy(source, transform)

	return Flatten(many)
}

func MapTransformBy[K comparable, T1, T2 any](source map[K]T1, transform func(T1) T2) map[K]T2 {
	var result = make(map[K]T2, len(source))
	for k, v := range source {
		result[k] = transform(v)
	}

	return result
}

func MapToSlice[K comparable, T1 any, T2 any](source map[K]T1, transform func(key K, value T1) T2) []T2 {
	var result = make([]T2, 0, len(source))
	for key, value := range source {
		result = append(result, transform(key, value))
	}

	return result
}

func SliceToMap[K comparable, T any](source []T, keyFunc func(T) K) map[K]T {
	var result = make(map[K]T, len(source))
	for _, v := range source {
		result[keyFunc(v)] = v
	}

	return result
}

func Flatten[T any](source [][]T) []T {
	var size int
	for _, v := range source {
		size += len(v)
	}

	var result = make([]T, 0, size)
	for _, v := range source {
		result = append(result, v...)
	}

	return result
}

// Duplicates returns elements that occur in collection more than once
func Duplicates[T comparable](source []T) []T {
	var asMap = make(map[T]struct{}, len(source))
	var result = make([]T, 0, len(source))

	Each(source, func(v T) {
		if _, ok := asMap[v]; ok {
			result = append(result, v)
		}

		asMap[v] = struct{}{}
	})

	return Distinct(result)
}
