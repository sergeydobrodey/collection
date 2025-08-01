package collection

// Aggregate aggregates the elements of the slice into a single value using a user-defined aggregator function.
func Aggregate[S ~[]T, T, K any](source S, aggregator func(K, T) K) K {
	var result K

	for _, v := range source {
		result = aggregator(result, v)
	}

	return result
}
