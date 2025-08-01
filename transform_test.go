package collection_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"maps"
	"slices"

	"github.com/sergeydobrodey/collection"
)

func TestDuplicates(t *testing.T) {
	var (
		cases = []struct {
			source []string
			want   []string
		}{
			{source: []string{}, want: []string{}},
			{source: []string{"a", "b", "a", "a"}, want: []string{"a"}},
		}
	)

	for _, tc := range cases {
		got := collection.Duplicates(tc.source)

		if !slices.Equal(got, tc.want) {
			t.Errorf("Duplicates(%v) = %v; want %v", tc.source, got, tc.want)
		}
	}
}

func TestTransformBy(t *testing.T) {
	cases := []struct {
		name      string
		source    []int
		transform func(v int) string
		want      []string
	}{
		{
			name:      "integers to strings",
			source:    []int{1, 2, 3, 4},
			transform: strconv.Itoa,
			want:      []string{"1", "2", "3", "4"},
		},
		{
			name:      "integers to strings of their squares",
			source:    []int{1, 2, 3, 4},
			transform: func(n int) string { return strconv.Itoa(n * n) },
			want:      []string{"1", "4", "9", "16"},
		},
		{
			name:      "an empty slice",
			source:    []int{},
			transform: strconv.Itoa,
			want:      []string{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.TransformBy(tc.source, tc.transform)

			if !slices.Equal(got, tc.want) {
				t.Errorf("TransformBy(%v) = %v; want %v", tc.source, got, tc.want)
			}
		})
	}
}

func TestTransformManyBy(t *testing.T) {
	cases := []struct {
		name      string
		source    []string
		transform func(string) []int
		want      []int
	}{
		{
			name:   "Transform strings to lengths of words",
			source: []string{"apple", "banana", "cherry"},
			transform: func(s string) []int {
				return []int{len(s)}
			},
			want: []int{5, 6, 6},
		},
		{
			name:   "Transform strings to ASCII values of characters",
			source: []string{"ab", "cd"},
			transform: func(s string) []int {
				result := make([]int, len(s))
				for i, char := range s {
					result[i] = int(char)
				}
				return result
			},
			want: []int{97, 98, 99, 100}, // ASCII values of 'a', 'b', 'c', 'd'
		},
		{
			name:   "Empty list",
			source: []string{},
			transform: func(s string) []int {
				return []int{len(s)}
			},
			want: []int{},
		},
		{
			name:   "Transform strings to position of characters in alphabet",
			source: []string{"abc", "xyz"},
			transform: func(s string) []int {
				result := make([]int, len(s))
				for i, char := range s {
					result[i] = int(char - 'a' + 1)
				}
				return result
			},
			want: []int{1, 2, 3, 24, 25, 26}, // Positions of 'a', 'b', 'c', 'x', 'y', 'z'
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.TransformManyBy(tc.source, tc.transform)

			if !slices.Equal(got, tc.want) {
				t.Errorf("TransformManyBy(%v) = %v; want %v", tc.source, got, tc.want)
			}
		})
	}
}

func TestAsyncTryTransformBy(t *testing.T) {
	someError := fmt.Errorf("some error")

	cases := []struct {
		name         string
		source       []string
		transform    func(context.Context, string) (int, error)
		want         []int
		wantErr      error
		cancelParent bool
	}{
		{
			name:   "Successful transformation",
			source: []string{"1", "2", "3"},
			transform: func(ctx context.Context, s string) (int, error) {
				return len(s), nil
			},
			want:    []int{1, 1, 1},
			wantErr: nil,
		},
		{
			name:   "Error scenario",
			source: []string{"1", "2", "3"},
			transform: func(ctx context.Context, s string) (int, error) {
				if s == "2" {
					return -1, someError
				}

				return 0, nil
			},
			want:    nil,
			wantErr: someError,
		},
		{
			name:   "Cancellation of asynchronous transformations",
			source: []string{"1", "2", "3"},
			transform: func(ctx context.Context, s string) (int, error) {
				if s == "3" {
					var timeout, cancel = context.WithTimeout(ctx, 500*time.Millisecond)
					defer cancel()

					<-timeout.Done()

					return -1, timeout.Err()
				}

				return strconv.Atoi(s)
			},
			want:         nil,
			wantErr:      context.Canceled,
			cancelParent: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			if tc.cancelParent {
				cancel()
			}

			got, err := collection.AsyncTryTransformBy(ctx, tc.source, tc.transform)

			if tc.wantErr == nil && err != nil {
				t.Errorf("AsyncTryTransformBy(%v) = %v; want %v", tc.source, err, tc.wantErr)
			}

			if tc.wantErr != nil && (err == nil || tc.wantErr.Error() != err.Error()) {
				t.Errorf("AsyncTryTransformBy(%v) = %v; want %v", tc.source, err, tc.wantErr)
			}

			slices.Sort(got)

			if !slices.Equal(got, tc.want) {
				t.Errorf("AsyncTryTransformBy(%v) = %v; want %v", tc.source, got, tc.want)
			}
		})
	}
}

func TestTryTransformBy(t *testing.T) {
	cases := []struct {
		name      string
		source    []string
		transform func(string) (int, error)
		want      []int
		wantErr   error
	}{
		{
			name:      "Normal transformation",
			source:    []string{"1", "2", "3"},
			transform: strconv.Atoi,
			want:      []int{1, 2, 3},
			wantErr:   nil,
		},
		{
			name:      "Transformation with an error",
			source:    []string{"1", "a", "3"},
			transform: strconv.Atoi,
			want:      nil,
			wantErr:   fmt.Errorf("strconv.Atoi: parsing \"a\": invalid syntax"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := collection.TryTransformBy(tc.source, tc.transform)

			if tc.wantErr == nil && err != nil {
				t.Errorf("TryTransformBy(%v) = %v; want %v", tc.source, err, tc.wantErr)
			}

			if tc.wantErr != nil && (err == nil || tc.wantErr.Error() != err.Error()) {
				t.Errorf("TryTransformBy(%v) = %v; want %v", tc.source, err, tc.wantErr)
			}

			if !slices.Equal(got, tc.want) {
				t.Errorf("TryTransformBy(%v) = %v; want %v", tc.source, got, tc.want)
			}
		})
	}
}

func TestMapTransformBy(t *testing.T) {
	cases := []struct {
		name      string
		source    map[string]int
		transform func(int) float64
		want      map[string]float64
	}{
		{
			name:      "Normal transformation",
			source:    map[string]int{"a": 1, "b": 2, "c": 3},
			transform: func(i int) float64 { return float64(i) + 0.5 },
			want:      map[string]float64{"a": 1.5, "b": 2.5, "c": 3.5},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.MapTransformBy(tc.source, tc.transform)

			if !maps.Equal(got, tc.want) {
				t.Errorf("MapTransformBy(%v) = %v; want %v", tc.source, got, tc.want)
			}
		})
	}
}

func TestTryMapTransformBy(t *testing.T) {
	transformError := fmt.Errorf("transform error")

	cases := []struct {
		name      string
		source    map[string]int
		transform func(int) (float64, error)
		want      map[string]float64
		wantErr   error
	}{
		{
			name:   "Transform all values successfully",
			source: map[string]int{"a": 1, "b": 2, "c": 3},
			transform: func(val int) (float64, error) {
				return float64(val) * 2.0, nil
			},
			want:    map[string]float64{"a": 2.0, "b": 4.0, "c": 6.0},
			wantErr: nil,
		},
		{
			name:   "Error during transformation",
			source: map[string]int{"a": 1, "b": 2, "c": 3},
			transform: func(val int) (float64, error) {
				return 0.0, transformError
			},
			want:    nil,
			wantErr: transformError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := collection.TryMapTransformBy(tc.source, tc.transform)

			if tc.wantErr == nil && err != nil {
				t.Errorf("TryMapTransformBy(%v) = %v; want %v", tc.source, err, tc.wantErr)
			}

			if tc.wantErr != nil && (err == nil || tc.wantErr.Error() != err.Error()) {
				t.Errorf("TryMapTransformBy(%v) = %v; want %v", tc.source, err, tc.wantErr)
			}

			if !maps.Equal(got, tc.want) {
				t.Errorf("TryMapTransformBy(%v) = %v; want %v", tc.source, got, tc.want)
			}
		})
	}
}

func TestMapToSlice(t *testing.T) {
	cases := []struct {
		name      string
		source    map[string]int
		transform func(key string, value int) float64
		want      []float64
	}{
		{
			name:   "Empty map",
			source: make(map[string]int),
			transform: func(key string, value int) float64 {
				return float64(value) * 1.5
			},
			want: []float64{},
		},
		{
			name:   "Map with positive values",
			source: map[string]int{"a": 1, "b": 2, "c": 3},
			transform: func(key string, value int) float64 {
				return float64(value) * 1.5
			},
			want: []float64{1.5, 3, 4.5},
		},
		{
			name:   "Map with negative values",
			source: map[string]int{"a": -1, "b": -2, "c": -3},
			transform: func(key string, value int) float64 {
				return float64(value) * 1.5
			},
			want: []float64{-1.5, -3, -4.5},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.MapToSlice(tc.source, tc.transform)

			slices.Sort(got)
			slices.Sort(tc.want)

			if !slices.Equal(got, tc.want) {
				t.Errorf("MapToSlice(%v) = %v; want %v", tc.source, got, tc.want)
			}
		})
	}
}

func TestSliceToMap(t *testing.T) {
	testCases := []struct {
		name    string
		source  []string
		keyFunc func(string) int
		want    map[int]string
	}{
		{
			name:    "Empty slice",
			source:  []string{},
			keyFunc: func(s string) int { return len(s) },
			want:    map[int]string{},
		},
		{
			name:    "Non-empty slice",
			source:  []string{"apple", "banana", "cherry"},
			keyFunc: func(s string) int { return len(s) },
			want:    map[int]string{5: "apple", 6: "cherry"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.SliceToMap(tc.source, tc.keyFunc)

			if !maps.Equal(got, tc.want) {
				t.Errorf("SliceToMap(%v) = %v; want %v", tc.source, got, tc.want)
			}
		})
	}
}

func TestAsyncTransformBy(t *testing.T) {
	cases := []struct {
		name      string
		source    []int
		transform func(v int) string
		want      []string
	}{
		{
			name:      "integers to strings",
			source:    []int{1, 2, 3, 4},
			transform: strconv.Itoa,
			want:      []string{"1", "2", "3", "4"},
		},
		{
			name:      "integers to strings of their squares",
			source:    []int{1, 2, 3, 4},
			transform: func(n int) string { return strconv.Itoa(n * n) },
			want:      []string{"1", "4", "9", "16"},
		},
		{
			name:      "an empty slice",
			source:    []int{},
			transform: strconv.Itoa,
			want:      []string{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.AsyncTransformBy(tc.source, tc.transform)

			if !slices.Equal(got, tc.want) {
				t.Errorf("TransformBy(%v) = %v; want %v", tc.source, got, tc.want)
			}
		})
	}
}

func TestChunkBy(t *testing.T) {
	cases := []struct {
		name   string
		source []int
		size   int
		want   [][]int
	}{
		{
			name:   "evenly divisible chunks",
			source: []int{1, 2, 3, 4, 5, 6},
			size:   2,
			want:   [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name:   "remainder chunk",
			source: []int{1, 2, 3, 4, 5},
			size:   2,
			want:   [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name:   "chunk size larger than source",
			source: []int{1, 2, 3},
			size:   5,
			want:   [][]int{{1, 2, 3}},
		},
		{
			name:   "chunk size equal to source length",
			source: []int{1, 2, 3},
			size:   3,
			want:   [][]int{{1, 2, 3}},
		},
		{
			name:   "empty source slice",
			source: []int{},
			size:   3,
			want:   nil,
		},
		{
			name:   "zero chunk size",
			source: []int{1, 2, 3},
			size:   0,
			want:   nil,
		},
		{
			name:   "negative chunk size",
			source: []int{1, 2, 3},
			size:   -1,
			want:   nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.ChunkBy(tc.source, tc.size)

			if !slices.EqualFunc(got, tc.want, slices.Equal[[]int]) {
				t.Errorf("ChunkBy(%v, %d) = %v; want %v", tc.source, tc.size, got, tc.want)
			}
		})
	}
}
