package collection

import (
	"golang.org/x/exp/constraints"
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
