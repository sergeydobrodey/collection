package collection_test

import (
	"maps"
	"slices"
	"testing"

	"github.com/sergeydobrodey/collection"
)

func TestCopy(t *testing.T) {
	cases := []struct {
		name   string
		source []int
	}{
		{name: "empty slice", source: nil},
		{name: "slice 6 elems", source: []int{4, 8, 15, 16, 23, 42}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.Copy(tc.source)

			if !slices.Equal(got, tc.source) {
				t.Errorf("Copy(%v) = %v; want %v", tc.source, got, tc.source)
			}
		})
	}
}

func TestClone(t *testing.T) {
	cases := []struct {
		name   string
		source []int
	}{
		{name: "empty slice", source: []int{}},
		{name: "nil slice", source: nil},
		{name: "slice with elements", source: []int{1, 2, 3, 4, 5}},
		{name: "slice with duplicates", source: []int{1, 1, 2, 2, 3}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.Clone(tc.source)

			if !slices.Equal(got, tc.source) {
				t.Errorf("Clone(%v) = %v; want %v", tc.source, got, tc.source)
			}

			// Test that modifying the original doesn't affect the clone
			if len(tc.source) > 0 {
				original := make([]int, len(tc.source))
				copy(original, tc.source)
				tc.source[0] = 999

				if slices.Equal(got, tc.source) {
					t.Errorf("Clone should be independent of original, but they are equal after modification")
				}

				if !slices.Equal(got, original) {
					t.Errorf("Clone should equal the original before modification, got %v, want %v", got, original)
				}
			}
		})
	}
}

func TestMapClone(t *testing.T) {
	cases := []struct {
		name   string
		source map[string]int
	}{
		{name: "empty map", source: map[string]int{}},
		{name: "nil map", source: nil},
		{name: "map with elements", source: map[string]int{"a": 1, "b": 2, "c": 3}},
		{name: "map with duplicates", source: map[string]int{"a": 1, "b": 1, "c": 2}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.MapClone(tc.source)

			if !maps.Equal(got, tc.source) {
				t.Errorf("MapClone(%v) = %v; want %v", tc.source, got, tc.source)
			}

			// Test that modifying the original doesn't affect the clone
			if len(tc.source) > 0 {
				original := make(map[string]int, len(tc.source))
				for k, v := range tc.source {
					original[k] = v
				}

				// Modify the original
				for k := range tc.source {
					tc.source[k] = 999
					break // Only modify one key
				}

				if maps.Equal(got, tc.source) {
					t.Errorf("MapClone should be independent of original, but they are equal after modification")
				}

				if !maps.Equal(got, original) {
					t.Errorf("MapClone should equal the original before modification, got %v, want %v", got, original)
				}
			}
		})
	}
}
