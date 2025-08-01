package collection

import (
	"slices"

	"golang.org/x/exp/constraints"
	"maps"
)

// Min returns the smaller of x or y.
func Min[T constraints.Ordered](l T, r T) T {
	if l <= r {
		return l
	}

	return r
}

// Max returns the larger of x or y.
func Max[T constraints.Ordered](l T, r T) T {
	if l >= r {
		return l
	}

	return r
}

// MinOf returns the smallest value among the provided elements or zero value
func MinOf[T constraints.Ordered](elements ...T) T {
	if len(elements) == 0 {
		var zero T
		return zero
	}

	min := elements[0]
	for _, v := range elements {
		if v < min {
			min = v
		}
	}

	return min
}

// MaxOf returns the largest value among the provided elements or zero value
func MaxOf[T constraints.Ordered](elements ...T) T {
	if len(elements) == 0 {
		var zero T
		return zero
	}

	max := elements[0]
	for _, v := range elements {
		if v > max {
			max = v
		}
	}

	return max
}

// Equal is equal to slices.Equal
func Equal[S ~[]T, T comparable](s1, s2 S) bool {
	return slices.Equal(s1, s2)
}

// MapEqual is equal to maps.Equal
func MapEqual[M1, M2 ~map[K]V, K, V comparable](m1 M1, m2 M2) bool {
	return maps.Equal(m1, m2)
}

// EqualFunc is equal to slices.EqualFunc
func EqualFunc[S1 ~[]T1, S2 ~[]T2, T1, T2 any](s1 S1, s2 S2, eq func(T1, T2) bool) bool {
	return slices.EqualFunc(s1, s2, eq)
}

// MapEqualFunc is equal to maps.EqualFunc
func MapEqualFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](m1 M1, m2 M2, eq func(V1, V2) bool) bool {
	return maps.EqualFunc(m1, m2, eq)
}
