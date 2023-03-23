package collection

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](l T, r T) T {
	if l <= r {
		return l
	}

	return r
}

func Max[T constraints.Ordered](l T, r T) T {
	if l >= r {
		return l
	}

	return r
}

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
