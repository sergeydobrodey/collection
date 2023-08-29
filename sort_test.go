package collection_test

import (
	"reflect"
	"testing"

	"github.com/sergeydobrodey/collection"
)

func TestSort(t *testing.T) {
	cases := []struct {
		source []int
		want   []int
	}{
		{[]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}},
	}

	for _, tc := range cases {
		collection.Sort(tc.source)
		if !reflect.DeepEqual(tc.source, tc.want) {
			t.Errorf("Sort(%v) = %v, want %v", tc.source, tc.source, tc.want)
		}
	}
}

func TestSortBy(t *testing.T) {
	cases := []struct {
		source []string
		less   func(l, r string) bool
		want   []string
	}{
		{[]string{"hello", "world", "foo", "bar", "baz"}, func(l, r string) bool { return l < r }, []string{"bar", "baz", "foo", "hello", "world"}},
		{[]string{"aaa", "bb", "c", "dddddd"}, func(l, r string) bool { return len(l) < len(r) }, []string{"c", "bb", "aaa", "dddddd"}},
	}

	for _, tc := range cases {
		collection.SortBy(tc.source, tc.less)
		if !reflect.DeepEqual(tc.source, tc.want) {
			t.Errorf("SortBy(%v) = %v, want %v", tc.source, tc.source, tc.want)
		}
	}
}

func TestReverse(t *testing.T) {
	cases := []struct {
		source []int
		want   []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}},
	}

	for _, tc := range cases {
		collection.Reverse(tc.source)
		if !reflect.DeepEqual(tc.source, tc.want) {
			t.Errorf("Reverse(%v) = %v, want %v", tc.source, tc.source, tc.want)
		}
	}
}
