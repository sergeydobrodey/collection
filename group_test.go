package collection_test

import (
	"reflect"
	"testing"

	"github.com/sergeydobrodey/collection"
)

func TestGroupBy(t *testing.T) {
	isEven := func(x int) bool { return x%2 == 0 }
	isOdd := func(x int) bool { return x%2 != 0 }

	cases := []struct {
		source  []int
		keyFunc func(i int) bool
		want    []int
	}{
		{[]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, isOdd, []int{1, 1, 3, 3, 5, 5, 5, 9}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, isEven, []int{0, 2, 4, 6, 8}},
	}

	for _, tc := range cases {
		var groups = collection.GroupBy(tc.source, tc.keyFunc)
		var result = groups[true]
		collection.Sort(result)

		if !reflect.DeepEqual(result, tc.want) {
			t.Errorf("GroupBy = %v, want %v", tc.source, tc.want)
		}
	}
}
