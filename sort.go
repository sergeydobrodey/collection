package collection

import (
	"slices"
	"sort"

	"golang.org/x/exp/constraints"
)

// Sort sorts the source slice of type T in ascending order.
func Sort[S ~[]T, T constraints.Ordered](source S) {
	slices.Sort(source)
}

// SortBy sorts the source slice of type T according to the less function provided.
func SortBy[S ~[]T, T any](source S, less func(l T, r T) bool) {
	sort.Slice(source, func(i, j int) bool {
		return less(source[i], source[j])
	})
}

// Reverse reverses the order of the elements in the source slice of type T.
func Reverse[S ~[]T, T any](source S) {
	for i, j := 0, len(source)-1; i < j; i, j = i+1, j-1 {
		source[i], source[j] = source[j], source[i]
	}
}
