package collection

type Filter[T any] func(T) bool

func InFilter[T comparable](source []T) Filter[T] {
	var set = make(map[T]struct{}, len(source))
	for _, v := range source {
		set[v] = struct{}{}
	}

	return func(item T) bool {
		_, ok := set[item]
		return ok
	}
}

func FilterBy[T any](source []T, filter Filter[T]) []T {
	var result = make([]T, 0, len(source))
	for _, item := range source {
		if filter(item) {
			result = append(result, item)
		}
	}

	return result
}

func MapFilterBy[K comparable, T any](source map[K]T, filter func(key K, value T) bool) map[K]T {
	var result = make(map[K]T, len(source))
	for key, value := range source {
		if filter(key, value) {
			result[key] = value
		}
	}

	return result
}

func Distinct[T comparable](source []T) []T {
	var set = make(map[T]struct{}, len(source))
	for _, v := range source {
		set[v] = struct{}{}
	}

	return MapKeys(set)
}

// Difference finds a set difference between a and b
// (values that are in a but not in b or a-b).
func Difference[T comparable](a []T, b []T) []T {
	var (
		ok   bool
		bMap = SliceToMap(b, func(t T) T { return t })
	)

	return Distinct(FilterBy(a, func(t T) bool {
		_, ok = bMap[t]
		// we need a elements that are not in B
		return !ok
	}))
}
