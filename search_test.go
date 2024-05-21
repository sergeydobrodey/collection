package collection_test

import (
	"testing"

	"github.com/sergeydobrodey/collection"
)

func TestContains(t *testing.T) {
	cases := []struct {
		name   string
		source []int
		item   int
		want   bool
	}{
		{"Item present", []int{1, 2, 3, 4}, 3, true},
		{"Item absent", []int{1, 2, 3, 4}, 5, false},
		{"Empty slice", []int{}, 1, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.Contains(tc.source, tc.item)

			if got != tc.want {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}

func TestMapContains(t *testing.T) {
	cases := []struct {
		name   string
		source map[int]string
		item   int
		want   bool
	}{
		{"Key present", map[int]string{1: "a", 2: "b", 3: "c"}, 2, true},
		{"Key absent", map[int]string{1: "a", 2: "b", 3: "c"}, 4, false},
		{"Empty map", map[int]string{}, 1, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.MapContains(tc.source, tc.item)

			if got != tc.want {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}

func TestAny(t *testing.T) {
	cases := []struct {
		name      string
		source    []int
		predicate func(int) bool
		want      bool
	}{
		{"At least one even number", []int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 }, true},
		{"No even numbers", []int{1, 3, 5, 7}, func(n int) bool { return n%2 == 0 }, false},
		{"Empty slice", []int{}, func(n int) bool { return n%2 == 0 }, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.Any(tc.source, tc.predicate)

			if got != tc.want {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}

func TestAll(t *testing.T) {
	cases := []struct {
		name      string
		source    []int
		predicate func(int) bool
		want      bool
	}{
		{"All even numbers", []int{2, 4, 6, 8}, func(n int) bool { return n%2 == 0 }, true},
		{"Not all even numbers", []int{2, 4, 5, 8}, func(n int) bool { return n%2 == 0 }, false},
		{"Empty slice", []int{}, func(n int) bool { return n%2 == 0 }, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.All(tc.source, tc.predicate)

			if got != tc.want {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}
