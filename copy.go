package collection

// Copy returns copy of the slice
func Copy[T any](source []T) []T {
	var result = make([]T, len(source))
	copy(result, source)

	return result
}
