package collection_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/sergeydobrodey/collection"
)

func TestAggregate(t *testing.T) {
	type args struct {
		source     []int
		aggregator func(int, int) int
	}

	cases := []struct {
		name string
		args args
		want int
	}{
		{name: "sum", args: args{source: []int{1, 2, 3, 4, 5}, aggregator: func(s int, v int) int { return s + v }}, want: 15},
		{name: "sum even", args: args{source: []int{1, 2, 3, 4, 5}, aggregator: func(s int, v int) int {
			if v%2 != 0 {
				return s
			}

			return s + v
		}}, want: 6},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := collection.Aggregate(tc.args.source, tc.args.aggregator); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Aggregate() = %v, want %v", got, tc.want)
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
