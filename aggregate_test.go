package collection_test

import (
	"fmt"
	"testing"

	"github.com/sergeydobrodey/collection"
)

func TestAggregate(t *testing.T) {
	cases := []struct {
		name       string
		source     []int
		aggregator func(int, int) int
		want       int
	}{
		{name: "sum", source: []int{1, 2, 3, 4, 5}, aggregator: func(s int, v int) int { return s + v }, want: 15},
		{name: "sum even", source: []int{1, 2, 3, 4, 5}, aggregator: func(s int, v int) int {
			if v%2 != 0 {
				return s
			}

			return s + v
		}, want: 6},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.Aggregate(tc.source, tc.aggregator)

			if got != tc.want {
				t.Errorf("Aggregate(%v) = %v; want %v", tc.source, got, tc.want)
			}
		})
	}
}

// ExampleAggregate: Example function demonstrating the use of the Aggregate function.
func ExampleAggregate() {
	sum := func(s int, v int) int { return s + v }

	result := collection.Aggregate([]int{1, 2, 3, 4, 5}, sum)
	fmt.Println(result)
	// Output: 15
}
