package collection

import (
	"reflect"
	"testing"
)

func TestDuplicates(t *testing.T) {
	var (
		cases = []struct {
			want []string
			a    []string
		}{
			{want: []string{}, a: []string{}},
			{want: []string{"a"}, a: []string{"a", "b", "a", "a"}},
		}
		res []string
	)

	for idx, testCase := range cases {
		res = Duplicates(testCase.a)

		if !reflect.DeepEqual(res, testCase.want) {
			t.Fatalf("%v != %v, test case: %d", res, testCase.want, idx)
		}
	}
}
