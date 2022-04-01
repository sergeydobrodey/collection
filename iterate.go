package collection

func Each[T any](source []T, do func(T)) {
	for _, v := range source {
		do(v)
	}
}

func MapEach[K comparable, T any](source map[K]T, do func(key K, value T)) {
	for k, v := range source {
		do(k, v)
	}
}
