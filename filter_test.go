package collection

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	var (
		cases = []struct {
			want []string
			a    []string
			b    []string
		}{
			{want: []string{"a"}, a: []string{"a", "b"}, b: []string{"b"}},
			{want: []string{}, a: []string{"a", "b"}, b: []string{"b", "a"}},
			{want: []string{"a"}, a: []string{"a", "a", "b"}, b: []string{"b"}},
			{want: []string{}, a: []string{}},
		}
		res []string
	)

	for idx, testCase := range cases {
		res = Difference(testCase.a, testCase.b)

		if !reflect.DeepEqual(res, testCase.want) {
			t.Fatalf("%v != %v, test case: %d", res, testCase.want, idx)
		}
	}
}

func TestDistinct(t *testing.T) {
	var (
		cases = []struct {
			want []string
			a    []string
		}{
			{want: []string{}, a: []string{}},
			{want: []string{"a", "b"}, a: []string{"a", "b", "a"}},
			{want: []string{"a", "b"}, a: []string{"a", "b", "a", "b"}},
			{want: []string{"a", "b", "c"}, a: []string{"a", "b", "c"}},
		}
		res []string
	)

	for idx, testCase := range cases {
		res = Distinct(testCase.a)
		if !equalSet(res, testCase.want) {
			t.Fatalf("%v != %v, test case: %d", res, testCase.want, idx)
		}
	}
}
