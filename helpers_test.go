package collection

import (
	"testing"
)

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
