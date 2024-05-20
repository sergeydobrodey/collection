package collection_test

import (
	"testing"

	"github.com/sergeydobrodey/collection"
)

func TestEach(t *testing.T) {
	cases := []struct {
		name   string
		source []int
		do     func(*int) func(v int)
		want   int
	}{
		{name: "sum", source: []int{1, 2, 3, 4, 5}, do: func(i *int) func(v int) {
			return func(v int) { *i += v }
		}, want: 15 + 1},
		{name: "mul", source: []int{-8, 1, 2, 3}, do: func(i *int) func(v int) {
			return func(v int) { *i *= v }
		}, want: -48 * 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := 1

			collection.Each(tc.source, tc.do(&result))
			if result != tc.want {
				t.Errorf("Each = %d; want %d", result, tc.want)
			}
		})
	}
}

func TestMapEach(t *testing.T) {
	cases := []struct {
		name   string
		source map[string]int
		do     func(*int) func(k string, v int)
		want   int
	}{
		{name: "sum", source: map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}, do: func(i *int) func(k string, v int) {
			return func(k string, v int) { *i += v }
		}, want: 15 + 2},
		{name: "mul", source: map[string]int{"-8": -8, "1": 1, "2": 2, "3": 3}, do: func(i *int) func(k string, v int) {
			return func(k string, v int) { *i *= v }
		}, want: -48 * 2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := 2

			collection.MapEach(tc.source, tc.do(&result))
			if result != tc.want {
				t.Errorf("MapEach = %d; want %d", result, tc.want)
			}
		})
	}
}
