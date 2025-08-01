package collection_test

import (
	"testing"

	"github.com/sergeydobrodey/collection"
)

func TestMin(t *testing.T) {
	cases := []struct {
		name string
		l    int
		r    int
		want int
	}{
		{name: "positive numbers", l: 5, r: 10, want: 5},
		{name: "negative numbers", l: -5, r: -10, want: -10},
		{name: "same numbers", l: 10, r: 10, want: 10},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.Min(tc.l, tc.r)

			if got != tc.want {
				t.Errorf("Min(%v, %v) = %v; want %v", tc.l, tc.r, got, tc.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	cases := []struct {
		name string
		l    int
		r    int
		want int
	}{
		{name: "positive numbers", l: 5, r: 10, want: 10},
		{name: "negative numbers", l: -5, r: -10, want: -5},
		{name: "same numbers", l: 10, r: 10, want: 10},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.Max(tc.l, tc.r)

			if got != tc.want {
				t.Errorf("Max(%v, %v) = %v; want %v", tc.l, tc.r, got, tc.want)
			}
		})
	}
}

func TestMinOf(t *testing.T) {
	cases := []struct {
		name   string
		source []int
		want   int
	}{
		{name: "positive numbers", source: []int{5, 10, 3}, want: 3},
		{name: "negative numbers", source: []int{-5, -10, -3}, want: -10},
		{name: "same numbers", source: []int{10, 10, 10}, want: 10},
		{name: "empty slice", source: []int{}, want: 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.MinOf(tc.source...)

			if got != tc.want {
				t.Errorf("MinOf(%v) = %v; want %v", tc.source, got, tc.want)
			}
		})
	}
}

func TestMaxOf(t *testing.T) {
	cases := []struct {
		name   string
		source []int
		want   int
	}{
		{name: "positive numbers", source: []int{5, 10, 3}, want: 10},
		{name: "negative numbers", source: []int{-5, -10, -3}, want: -3},
		{name: "same numbers", source: []int{10, 10, 10}, want: 10},
		{name: "empty slice", source: []int{}, want: 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.MaxOf(tc.source...)

			if got != tc.want {
				t.Errorf("MaxOf(%v) = %v; want %v", tc.source, got, tc.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	cases := []struct {
		name string
		l    []int
		r    []int
		want bool
	}{
		{name: "equal", l: []int{5, 10, 3}, r: []int{5, 10, 3}, want: true},
		{name: "wrong order", l: []int{5, 10, 3}, r: []int{3, 10, 5}, want: false},
		{name: "empty", l: []int{}, r: nil, want: true},
		{name: "not equal", l: []int{1, 2, 3, 5, 10, 3}, r: []int{5, 10, 3}, want: false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.Equal(tc.l, tc.r)

			if got != tc.want {
				t.Errorf("Equal(%v, %v) = %v; want %v", tc.l, tc.r, got, tc.want)
			}
		})
	}
}

func TestMapEqual(t *testing.T) {
	cases := []struct {
		name string
		l    map[string]int
		r    map[string]int
		want bool
	}{
		{name: "equal", l: map[string]int{"5": 5, "10": 10, "3": 3}, r: map[string]int{"5": 5, "10": 10, "3": 3}, want: true},
		{name: "different order", l: map[string]int{"5": 5, "10": 10, "3": 3}, r: map[string]int{"3": 3, "5": 5, "10": 10}, want: true},
		{name: "empty", l: map[string]int{}, r: nil, want: true},
		{name: "not equal", l: map[string]int{"5": 5, "10": 10, "3": 3}, r: map[string]int{"5": 5, "12": 12}, want: false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.MapEqual(tc.l, tc.r)

			if got != tc.want {
				t.Errorf("MapEqual(%v, %v) = %v; want %v", tc.l, tc.r, got, tc.want)
			}
		})
	}
}

func TestEqualFunc(t *testing.T) {
	type Person struct {
		ID   int
		Name string
	}

	cases := []struct {
		name   string
		s1     []Person
		s2     []Person
		equals func(Person, Person) bool
		want   bool
	}{
		{
			name: "equal by ID",
			s1:   []Person{{1, "Alice"}, {2, "Bob"}},
			s2:   []Person{{1, "Alice"}, {2, "Bob"}},
			equals: func(a, b Person) bool {
				return a.ID == b.ID
			},
			want: true,
		},
		{
			name: "equal by name",
			s1:   []Person{{1, "Alice"}, {2, "Bob"}},
			s2:   []Person{{3, "Alice"}, {4, "Bob"}},
			equals: func(a, b Person) bool {
				return a.Name == b.Name
			},
			want: true,
		},
		{
			name: "different lengths",
			s1:   []Person{{1, "Alice"}, {2, "Bob"}},
			s2:   []Person{{1, "Alice"}},
			equals: func(a, b Person) bool {
				return a.ID == b.ID
			},
			want: false,
		},
		{
			name: "not equal",
			s1:   []Person{{1, "Alice"}, {2, "Bob"}},
			s2:   []Person{{1, "Alice"}, {3, "Charlie"}},
			equals: func(a, b Person) bool {
				return a.ID == b.ID
			},
			want: false,
		},
		{
			name:   "empty slices",
			s1:     []Person{},
			s2:     []Person{},
			equals: func(a, b Person) bool { return a.ID == b.ID },
			want:   true,
		},
		{
			name:   "nil slices",
			s1:     nil,
			s2:     nil,
			equals: func(a, b Person) bool { return a.ID == b.ID },
			want:   true,
		},
		{
			name:   "one nil slice",
			s1:     []Person{{1, "Alice"}},
			s2:     nil,
			equals: func(a, b Person) bool { return a.ID == b.ID },
			want:   false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.EqualFunc(tc.s1, tc.s2, tc.equals)

			if got != tc.want {
				t.Errorf("EqualFunc(%v, %v, equals) = %v; want %v", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}

func TestMapEqualFunc(t *testing.T) {
	type Person struct {
		ID   int
		Name string
	}

	cases := []struct {
		name        string
		m1          map[int]Person
		m2          map[int]Person
		valueEquals func(Person, Person) bool
		want        bool
	}{
		{
			name: "equal by name",
			m1:   map[int]Person{1: {1, "Alice"}, 2: {2, "Bob"}},
			m2:   map[int]Person{1: {3, "Alice"}, 2: {4, "Bob"}},
			valueEquals: func(a, b Person) bool {
				return a.Name == b.Name
			},
			want: true,
		},
		{
			name: "equal by ID",
			m1:   map[int]Person{1: {1, "Alice"}, 2: {2, "Bob"}},
			m2:   map[int]Person{1: {1, "Alice"}, 2: {2, "Bob"}},
			valueEquals: func(a, b Person) bool {
				return a.ID == b.ID && a.Name == b.Name
			},
			want: true,
		},
		{
			name: "different sizes",
			m1:   map[int]Person{1: {1, "Alice"}, 2: {2, "Bob"}},
			m2:   map[int]Person{1: {1, "Alice"}},
			valueEquals: func(a, b Person) bool {
				return a.ID == b.ID
			},
			want: false,
		},
		{
			name: "not equal values",
			m1:   map[int]Person{1: {1, "Alice"}, 2: {2, "Bob"}},
			m2:   map[int]Person{1: {1, "Alice"}, 2: {2, "Charlie"}},
			valueEquals: func(a, b Person) bool {
				return a.Name == b.Name
			},
			want: false,
		},
		{
			name:        "empty maps",
			m1:          map[int]Person{},
			m2:          map[int]Person{},
			valueEquals: func(a, b Person) bool { return a.ID == b.ID },
			want:        true,
		},
		{
			name:        "nil maps",
			m1:          nil,
			m2:          nil,
			valueEquals: func(a, b Person) bool { return a.ID == b.ID },
			want:        true,
		},
		{
			name:        "one nil map",
			m1:          map[int]Person{1: {1, "Alice"}},
			m2:          nil,
			valueEquals: func(a, b Person) bool { return a.ID == b.ID },
			want:        false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.MapEqualFunc(tc.m1, tc.m2, tc.valueEquals)

			if got != tc.want {
				t.Errorf("MapEqualFunc(%v, %v, valueEquals) = %v; want %v", tc.m1, tc.m2, got, tc.want)
			}
		})
	}
}
