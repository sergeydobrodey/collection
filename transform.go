package collection

import (
	"context"
	"errors"
	"sync"
)

// TransformBy transform the source slice of type T to a new slice of type K using the provided transform function.
func TransformBy[T, K any](source []T, transform func(T) K) []K {
	var result = make([]K, len(source))
	for i, item := range source {
		result[i] = transform(item)
	}

	return result
}

// TransformManyBy transforms the source slice of type T to multiple slices of type K using the provided transform function.
func TransformManyBy[T, K any](source []T, transform func(T) []K) []K {
	var many = TransformBy(source, transform)

	return Flatten(many)
}

// TryTransformBy tries to transform the source slice of type T to a new slice of type K using the provided transform function.
func TryTransformBy[T, K any](source []T, transform func(T) (K, error)) ([]K, error) {
	var result = make([]K, len(source))
	for i, item := range source {
		var value, err = transform(item)
		if err != nil {
			return nil, err
		}

		result[i] = value
	}

	return result, nil
}

// MapTransformBy transform the values of the source map of type T1 to a new map of type T2 using the provided transform function.
func MapTransformBy[K comparable, T1, T2 any](source map[K]T1, transform func(T1) T2) map[K]T2 {
	var result = make(map[K]T2, len(source))
	for k, v := range source {
		result[k] = transform(v)
	}

	return result
}

// TryMapTransformBy attempts to transform the values of the source map of type T1 to a new map of type T2 using the provided transform function.
func TryMapTransformBy[K comparable, T1, T2 any](source map[K]T1, transform func(T1) (T2, error)) (map[K]T2, error) {
	var result = make(map[K]T2, len(source))
	for k, v := range source {
		var value, err = transform(v)
		if err != nil {
			return nil, err
		}

		result[k] = value
	}

	return result, nil
}

// MapToSlice convert the source map of type T1 to a slice of type T2 using the provided transform function on each key-value pair.
func MapToSlice[K comparable, T1 any, T2 any](source map[K]T1, transform func(key K, value T1) T2) []T2 {
	var result = make([]T2, 0, len(source))
	for key, value := range source {
		result = append(result, transform(key, value))
	}

	return result
}

// SliceToMap convert the source slice of type T to a new map of type T with keys generated by the provided keyFunc.
func SliceToMap[K comparable, T any](source []T, keyFunc func(T) K) map[K]T {
	var result = make(map[K]T, len(source))
	for _, v := range source {
		result[keyFunc(v)] = v
	}

	return result
}

// Flatten flattens a slice of slices into a single slice.
func Flatten[T any](source [][]T) []T {
	var size int
	for _, v := range source {
		size += len(v)
	}

	var result = make([]T, 0, size)
	for _, v := range source {
		result = append(result, v...)
	}

	return result
}

// Duplicates returns a new slice with all elements that appear more than once in the original slice.
func Duplicates[T comparable](source []T) []T {
	var visited = make(map[T]struct{}, len(source))
	var result = FilterBy(source, func(v T) bool {
		if _, ok := visited[v]; ok {
			return true
		}

		visited[v] = struct{}{}
		return false
	})

	return Distinct(result)
}

// AsyncTryTransformBy tries to async transform the source slice of type T to a new slice of type K using the provided transform function.
func AsyncTryTransformBy[T, K any](parent context.Context, source []T, transform func(context.Context, T) (K, error)) ([]K, error) {
	var (
		resultsCh   = make(chan Pair[K, error], len(source))
		ctx, cancel = context.WithCancel(parent)
	)

	defer cancel()

	var wg sync.WaitGroup
	wg.Add(len(source))

	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	for _, item := range source {
		go func(item T) {
			defer wg.Done()

			var response, err = transform(ctx, item)
			select {
			case resultsCh <- Pair[K, error]{First: response, Second: err}:
			case <-ctx.Done():
			}
		}(item)
	}

	var (
		result = make([]K, 0, len(source))
		errs   []error
	)

	for r := range resultsCh {
		if r.Second != nil {
			errs = append(errs, r.Second)

			cancel()
			continue
		}

		result = append(result, r.First)
	}

	if err := errors.Join(errs...); err != nil {
		return nil, err
	}

	return result, nil
}
