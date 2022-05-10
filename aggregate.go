package collection

func Aggregate[T, K any](source []T, aggregator func(K, T) K) K {
	var result K

	for _, v := range source {
		result = aggregator(result, v)
	}

	return result
}
