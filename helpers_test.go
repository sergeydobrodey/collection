package collection

import (
	"testing"
)

func TestEqualSet(t *testing.T) {
	var (
		cases = []struct {
			want bool
			a    []string
			b    []string
		}{
			{want: true, a: []string{}, b: []string{}},
			{want: false, a: []string{"a", "b"}, b: []string{"a"}},
			{want: true, a: []string{"a", "b", "a", "a"}, b: []string{"a", "b"}},
			{want: true, a: []string{"a", "b", "c"}, b: []string{"b", "a", "c"}},
		}
		res bool
	)

	for idx, testCase := range cases {
		res = equalSet(testCase.a, testCase.b)

		if res != testCase.want {
			t.Fatalf("%v != %v, test case: %d", res, testCase.want, idx)
		}
	}
}
