package collection

import (
	"maps"
	"slices"
)

// Copy is left for compatibility with previous versions.
// Deprecated: use Clone instead.
func Copy[S ~[]T, T any](s S) S {
	return Clone(s)
}

// Clone is equal to slices.Clone
func Clone[S ~[]T, T any](s S) S {
	return slices.Clone(s)
}

// MapClone is equal to maps.Clone
func MapClone[M ~map[K]V, K comparable, V any](m M) M {
	return maps.Clone(m)
}
