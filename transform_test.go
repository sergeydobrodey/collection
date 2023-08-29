package collection_test

import (
	"reflect"
	"testing"

	"github.com/sergeydobrodey/collection"
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

	for _, tc := range cases {
		res = collection.Duplicates(tc.a)

		if !reflect.DeepEqual(res, tc.want) {
			t.Errorf("%v != %v", res, tc.want)
		}
	}
}
