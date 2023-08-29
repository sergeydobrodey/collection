package collection_test

import (
	"testing"

	"github.com/sergeydobrodey/collection"
)

func TestCopy(t *testing.T) {
	cases := []struct {
		name   string
		source []int
	}{
		{name: "empty slice", source: nil},
		{name: "slice 6 elems", source: []int{4, 8, 15, 16, 23, 42}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.Copy(tc.source)
			if len(got) != len(tc.source) {
				t.Errorf("Copy(%v) = %v, want %v", tc.source, got, tc.source)
			}

			for i := range tc.source {
				if tc.source[i] != got[i] {
					t.Errorf("Copy(%v) = %v, want %v", tc.source, got, tc.source)
				}
			}
		})
	}
}
