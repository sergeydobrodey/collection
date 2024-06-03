package collection

type Filter[T any] func(T) bool

// InFilter returns a filter function that filters elements based on whether they are present or absent in the given slice.
func InFilter[T comparable](source []T, present bool) Filter[T] {
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
func FilterBy[T any](source []T, filter Filter[T]) []T {
	var result = make([]T, 0, len(source))
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
func Distinct[T comparable](source []T) []T {
	var (
		set    = make(map[T]struct{}, len(source))
		result = make([]T, 0, len(source))
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
func DistinctBy[T any](source []T, equals func(left T, right T) bool) []T {
	var result = make([]T, 0, len(source))

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
func Difference[T comparable](a []T, b []T) []T {
	return FilterBy(a, InFilter(b, false))
}

// SetIntersection finds a set intersection between a and b
// (unique values that are in a and in b).
func SetIntersection[T comparable](a []T, b []T) []T {
	return Distinct(FilterBy(a, InFilter(b, true)))
}

// Intersection finds an intersection between a and b
func Intersection[T comparable](a []T, b []T) []T {
	var set = make(map[T]int, len(b))
	for _, v := range b {
		set[v] += 1
	}

	var result = make([]T, 0, Min(len(a), len(b)))
	for _, v := range a {
		if len(set) == 0 {
			return result
		}

		if count, ok := set[v]; count > 0 && ok {
			result = append(result, v)
			set[v] -= 1
		}
	}

	return result
}
