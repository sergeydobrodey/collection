# Collections

This is a Go package that provides several utility functions for working with collections.

## Installation

To install this package, run the following command:

```Golang
go get github.com/sergeydobrodey/collection
```

## Usage

```Golang
import "github.com/sergeydobrodey/collection"
```

### Collection Functions

This package provides several functions for working with collections:

*   `Aggregate[T, K any](source []T, aggregator func(K, T) K) K`
*   `Any[T any](source []T, predicate func(T) bool) bool`
*   `Contains[T comparable](source []T, item T) bool`
*   `Difference[T comparable](a []T, b []T) []T`
*   `Distinct[T comparable](source []T) []T`
*   `Duplicates[T comparable](source []T) []T`
*   `Each[T any](source []T, do func(T))`
*   `FilterBy[T any](source []T, filter Filter[T]) []T`
*   `Filter[T any] func(T) bool`
*   `Flatten[T any](source [][]T) []T`
*   `GroupBy[T any, K comparable](source []T, keyFunc func(T) K) map[K][]T`
*   `InFilter[T comparable](source []T, present bool) Filter[T]`
*   `Intersection[T comparable](a []T, b []T) []T`
*   `MapContains[K comparable, T any](source map[K]T, item K) bool`
*   `MapEach[K comparable, T any](source map[K]T, do func(key K, value T))`
*   `MapFilterBy[K comparable, T any](source map[K]T, filter func(key K, value T) bool) map[K]T`
*   `MapKeys[K comparable, T any](source map[K]T) []K`
*   `MapToSlice[K comparable, T1 any, T2 any](source map[K]T1, transform func(key K, value T1) T2) []T2`
*   `MapTransformBy[K comparable, T1, T2 any](source map[K]T1, transform func(T1) T2) map[K]T2`
*   `MapValues[K comparable, T any](source map[K]T) []T`
*   `MaxOf[T constraints.Ordered](elements ...T) T`
*   `Max[T constraints.Ordered](l T, r T) T`
*   `MinOf[T constraints.Ordered](elements ...T) T`
*   `Min[T constraints.Ordered](l T, r T) T`
*   `Reverse[T any](source []T)`
*   `SliceToMap[K comparable, T any](source []T, keyFunc func(T) K) map[K]T`
*   `SortBy[T any](source []T, less func(l T, r T) bool)`
*   `Sort[T constraints.Ordered](source []T)`
*   `TransformBy[T, K any](source []T, transform func(T) K) []K`
*   `TransformManyBy[T, K any](source []T, transform func(T) []K) []K`


### SyncMap

The `SyncMap` type is a thread-safe map that can be accessed concurrently. It provides the following methods:

*   `CompareAndDelete(key K, old V) (deleted bool)`
*   `CompareAndSwap(key K, old V, new V) bool`
*   `Delete(key K)`
*   `Load(key K) (value V, ok bool)`
*   `LoadAndDelete(key K) (value V, loaded bool)`
*   `LoadOrStore(key K, value V) (actual V, loaded bool)`
*   `Range(f func(key K, value V) bool)`
*   `Store(key K, value V)`
*   `Swap(key K, value V) (previous V, loaded bool)`

