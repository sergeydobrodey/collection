package collection_test

import (
	"testing"
)

// equalSet checks if tw slices are equal sets or
// have the same unique unordered elements.
func equalSet[T comparable](a []T, b []T) bool {
	var (
		mapA = make(map[T]struct{}, len(a))
		mapB = make(map[T]struct{}, len(b))
		ok   bool
	)

	for _, v := range a {
		mapA[v] = struct{}{}
	}
	for _, v := range b {
		mapB[v] = struct{}{}
	}

	for k := range mapA {
		if _, ok = mapB[k]; !ok {
			return false
		}
	}

	for k := range mapB {
		if _, ok = mapA[k]; !ok {
			return false
		}
	}

	return true
}

func TestEqualSet(t *testing.T) {
	cases := []struct {
		name string
		want bool
		a    []string
		b    []string
	}{
		{name: "empty", want: true, a: []string{}, b: []string{}},
		{name: "a-b <-> a", want: false, a: []string{"a", "b"}, b: []string{"a"}},
		{name: "a-b-a-a <-> a-b", want: true, a: []string{"a", "b", "a", "a"}, b: []string{"a", "b"}},
		{name: "a-b-c <-> b-a-c", want: true, a: []string{"a", "b", "c"}, b: []string{"b", "a", "c"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := equalSet(tc.a, tc.b)

			if got != tc.want {
				t.Errorf("equalSet(%v,%v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
