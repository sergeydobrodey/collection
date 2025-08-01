package collection_test

import (
	"testing"

	"slices"

	"github.com/sergeydobrodey/collection"
)

func TestGroupBy(t *testing.T) {
	isEven := func(x int) bool { return x%2 == 0 }
	isOdd := func(x int) bool { return x%2 != 0 }

	cases := []struct {
		name    string
		source  []int
		keyFunc func(i int) bool
		want    []int
	}{
		{"is odd", []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, isOdd, []int{1, 1, 3, 3, 5, 5, 5, 9}},
		{"is even", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, isEven, []int{0, 2, 4, 6, 8}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.GroupBy(tc.source, tc.keyFunc)

			match := got[true]
			slices.Sort(match)

			if !slices.Equal(match, tc.want) {
				t.Errorf("GroupBy(%v) = %v; want %v", tc.source, match, tc.want)
			}
		})
	}
}
