package collection_test

import (
	"strings"
	"testing"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/sergeydobrodey/collection"
)

func TestDifference(t *testing.T) {
	cases := []struct {
		name string
		a    []string
		b    []string
		want []string
	}{
		{name: "simple", a: []string{"a", "b"}, b: []string{"b"}, want: []string{"a"}},
		{name: "different order", a: []string{"a", "b"}, b: []string{"b", "a"}, want: []string{}},
		{name: "complex", a: []string{"a", "a", "b"}, b: []string{"b"}, want: []string{"a", "a"}},
		{name: "empty", a: []string{}, want: []string{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.Difference(tc.a, tc.b)

			if !slices.Equal(got, tc.want) {
				t.Errorf("Difference(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestSetIntersection(t *testing.T) {
	cases := []struct {
		name string
		want []string
		a    []string
		b    []string
	}{
		{name: "empty", a: []string{}, b: []string{}, want: []string{}},
		{name: "one", a: []string{"a", "b"}, b: []string{"b"}, want: []string{"b"}},
		{name: "full intersection", a: []string{"a", "b"}, b: []string{"b", "a"}, want: []string{"a", "b"}},
		{name: "one-2", a: []string{"a", "a", "b"}, b: []string{"b"}, want: []string{"b"}},
		{name: "local intersection ", a: []string{"a", "a", "b"}, b: []string{"b", "b", "a", "a", "c"}, want: []string{"a", "b"}},
		{name: "local intersection-2", a: []string{"b", "b", "a", "a", "c"}, b: []string{"a", "a", "b"}, want: []string{"b", "a"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.SetIntersection(tc.a, tc.b)

			if !slices.Equal(got, tc.want) {
				t.Errorf("SetIntersection(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})

	}
}

func TestIntersection(t *testing.T) {
	cases := []struct {
		name string
		want []string
		a    []string
		b    []string
	}{
		{name: "empty", a: []string{}, b: []string{}, want: []string{}},
		{name: "one", a: []string{"a", "b"}, b: []string{"b"}, want: []string{"b"}},
		{name: "one to empty", a: []string{"a", "b"}, b: []string{}, want: []string{}},
		{name: "full intersection", a: []string{"a", "b"}, b: []string{"b", "a"}, want: []string{"a", "b"}},
		{name: "one-2", a: []string{"a", "a", "b"}, b: []string{"b"}, want: []string{"b"}},
		{name: "local intersection", a: []string{"a", "a", "b"}, b: []string{"b", "b", "a", "a", "c"}, want: []string{"a", "a", "b"}},
		{name: "local intersection-2", a: []string{"b", "b", "a", "a", "c"}, b: []string{"a", "a", "b"}, want: []string{"b", "a", "a"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.Intersection(tc.a, tc.b)

			if !slices.Equal(got, tc.want) {
				t.Errorf("Intersection(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})

	}
}

func TestDistinct(t *testing.T) {
	cases := []struct {
		name string
		a    []string
		want []string
	}{
		{name: "empty", a: []string{}, want: []string{}},
		{name: "one duplicate", a: []string{"a", "b", "a"}, want: []string{"a", "b"}},
		{name: "some duplicates", a: []string{"a", "b", "a", "b"}, want: []string{"a", "b"}},
		{name: "all uniq", a: []string{"a", "b", "c"}, want: []string{"a", "b", "c"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.Distinct(tc.a)

			if !slices.Equal(got, tc.want) {
				t.Errorf("Distinct(%v) = %v, want %v", tc.a, got, tc.want)
			}
		})
	}
}

func TestMapFilterBy(t *testing.T) {
	cases := []struct {
		name   string
		source map[string]int
		filter func(key string, value int) bool
		want   map[string]int
	}{
		{
			name:   "Filter values greater than 10",
			source: map[string]int{"a": 5, "b": 15, "c": 20},
			filter: func(key string, value int) bool { return value > 10 },
			want: map[string]int{
				"b": 15,
				"c": 20,
			},
		},
		{
			name:   "Filter keys starting with 'a'",
			source: map[string]int{"apple": 1, "banana": 2, "avocado": 3},
			filter: func(key string, value int) bool { return key[0] == 'a' },
			want: map[string]int{
				"apple":   1,
				"avocado": 3,
			},
		},
		{
			name:   "Filter an empty map",
			source: map[string]int{},
			filter: func(key string, value int) bool { return true },
			want:   map[string]int{},
		},
		{
			name:   "Filter keys containing 'e' and values greater than 10",
			source: map[string]int{"one": 1, "three": 13, "five": 15},
			filter: func(key string, value int) bool { return value > 10 && strings.ContainsRune(key, 'e') },
			want: map[string]int{
				"three": 13,
				"five":  15,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.MapFilterBy(tc.source, tc.filter)

			if !maps.Equal(got, tc.want) {
				t.Errorf("MapFilterBy(%v) = want %v, got %v", tc.source, tc.want, got)
			}
		})
	}
}

func TestDistinctBy(t *testing.T) {
	cases := []struct {
		name   string
		source []string
		equals func(left, right string) bool
		want   []string
	}{
		{
			name:   "Remove duplicate strings",
			source: []string{"apple", "banana", "apple", "cherry", "banana"},
			equals: func(left, right string) bool { return left == right },
			want: []string{
				"apple", "banana", "cherry",
			},
		},
		{
			name:   "All unique strings",
			source: []string{"apple", "banana", "cherry"},
			equals: func(left, right string) bool { return left == right },
			want:   []string{"apple", "banana", "cherry"},
		},
		{
			name:   "Empty slice",
			source: []string{},
			equals: func(left, right string) bool { return left == right },
			want:   []string{},
		},
		{
			name:   "Adjacent duplicates",
			source: []string{"apple", "apple", "banana", "banana", "cherry", "cherry"},
			equals: func(left, right string) bool { return left == right },
			want: []string{
				"apple", "banana", "cherry",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.DistinctBy(tc.source, tc.equals)

			if !slices.Equal(got, tc.want) {
				t.Errorf("DistinctBy(%v) = want %v, got %v", tc.source, tc.want, got)
			}
		})
	}
}
