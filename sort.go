package collection

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Ordered](source []T) {
	sort.Slice(source, func(i, j int) bool {
		return source[i] < source[j]
	})
}

func SortBy[T any](source []T, less func(l T, r T) bool) {
	sort.Slice(source, func(i, j int) bool {
		return less(source[i], source[j])
	})
}

func Reverse[T any](source []T) {
	for i, j := 0, len(source)-1; i < j; i, j = i+1, j-1 {
		source[i], source[j] = source[j], source[i]
	}
}
