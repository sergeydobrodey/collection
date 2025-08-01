package collection

type Filter[T any] func(T) bool

// InFilter returns a filter function that filters elements based on whether they are present or absent in the given slice.
func InFilter[S ~[]T, T comparable](source S, present bool) Filter[T] {
	var set = make(map[T]struct{}, len(source))
	for _, v := range source {
		set[v] = struct{}{}
	}

	return func(item T) bool {
		_, ok := set[item]
		return ok == present
	}
}

// FilterBy returns a new slice with only the elements that satisfy the given filter function.
func FilterBy[S ~[]T, T any](source S, filter Filter[T]) S {
	var result = make(S, 0, len(source))
	for _, item := range source {
		if filter(item) {
			result = append(result, item)
		}
	}

	return result
}

// MapFilterBy returns a new map with only the key-value pairs that satisfy the given filter function.
func MapFilterBy[K comparable, T any](source map[K]T, filter func(key K, value T) bool) map[K]T {
	var result = make(map[K]T, len(source))
	for key, value := range source {
		if filter(key, value) {
			result[key] = value
		}
	}

	return result
}

// Distinct returns a new slice with all duplicate elements removed.
func Distinct[S ~[]T, T comparable](source S) S {
	var (
		set    = make(map[T]struct{}, len(source))
		result = make(S, 0, len(source))
	)

	for _, v := range source {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

// DistinctBy returns a new slice with all duplicate elements removed.
func DistinctBy[S ~[]T, T any](source S, equals func(left T, right T) bool) S {
	var result = make(S, 0, len(source))

sourceLoop:
	for _, v := range source {
		for i, u := range result {
			if equals(v, u) {
				result[i] = v
				continue sourceLoop
			}
		}

		result = append(result, v)
	}

	return result
}

// Difference finds a set difference between a and b
// (values that are in a but not in b or a-b).
func Difference[S ~[]T, T comparable](a S, b S) S {
	return FilterBy(a, InFilter(b, false))
}

// Intersection finds a set intersection between a and b
// (unique values that are in a and in b).
func Intersection[S ~[]T, T comparable](a S, b S) S {
	return Distinct(FilterBy(a, InFilter(b, true)))
}
