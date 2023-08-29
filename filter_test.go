package collection

import (
	"reflect"
	"testing"
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
			got := Difference(tc.a, tc.b)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Difference(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
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
		{name: "full intersection", a: []string{"a", "b"}, b: []string{"b", "a"}, want: []string{"a", "b"}},
		{name: "one-2", a: []string{"a", "a", "b"}, b: []string{"b"}, want: []string{"b"}},
		{name: "local intersection ", a: []string{"a", "a", "b"}, b: []string{"b", "b", "a", "a", "c"}, want: []string{"a", "b"}},
		{name: "local intersection-2", a: []string{"b", "b", "a", "a", "c"}, b: []string{"a", "a", "b"}, want: []string{"b", "a"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Intersection(tc.a, tc.b)
			if !reflect.DeepEqual(got, tc.want) {
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
			got := Distinct(tc.a)
			if !equalSet(got, tc.want) {
				t.Errorf("Distinct(%v) = %v, want %v", tc.a, got, tc.want)
			}
		})
	}
}
