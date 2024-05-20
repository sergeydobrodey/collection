package collection

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Sort sorts the source slice of type T in ascending order.
func Sort[T constraints.Ordered](source []T) {
	slices.Sort(source)
}

// SortBy sorts the source slice of type T according to the less function provided.
func SortBy[T any](source []T, less func(l T, r T) bool) {
	slices.SortFunc(source, less)
}

// Reverse reverses the order of the elements in the source slice of type T.
func Reverse[T any](source []T) {
	for i, j := 0, len(source)-1; i < j; i, j = i+1, j-1 {
		source[i], source[j] = source[j], source[i]
	}
}
