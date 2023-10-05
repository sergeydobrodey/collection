# Collections

Collections is a Go package that provides an extensive collection of utility functions and a thread-safe map for working with collections. It includes a variety of functions such as filtering, mapping, sorting, and merging, as well as methods for accessing and modifying a map concurrently. Install the package with a simple `go get` command and import it into your project. Get started with the list of functions provided and streamline your collection-based workloads today!

## Installation

To install this package, run the following command:

```Golang
go get github.com/sergeydobrodey/collection
```

## Usage

```Golang
import "github.com/sergeydobrodey/collection"
```

### Documentation

Documentation is hosted at https://pkg.go.dev/github.com/sergeydobrodey/collection.

### Examples
#### Slice transformations
```golang
package main

import (
	"fmt"

	"github.com/sergeydobrodey/collection"
)

type User struct {
	id   uint
	name string
}

func (u User) ID() uint {
	return u.id
}

func (u User) Name() string {
	return u.name
}

func main() {
	var users = []User{{0, "Rob"}, {1, "Ken"}}

	var names = collection.TransformBy(users, User.Name)
	fmt.Println(names)
	// Output: [Rob Ken]

	var usersByID = collection.SliceToMap(users, User.ID)
	fmt.Println(usersByID)
	// Output: map[0:{0 Rob} 1:{1 Ken}]
}
```

#### Channels aggregation
```golang
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sergeydobrodey/collection"
)

type Worker interface {
	Run(context.Context)
	Done() <-chan struct{}
}

func NewWorker() Worker {
	return &worker{
		done: make(chan struct{}),
	}
}

type worker struct {
	done chan struct{}
}

func (w *worker) Done() <-chan struct{} {
	return w.done
}

func (w *worker) Run(ctx context.Context) {
	defer close(w.done)

	<-ctx.Done()
}

func main() {
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var worker1 = NewWorker()
	go worker1.Run(ctx)

	var worker2 = NewWorker()
	go worker2.Run(ctx)

	var allWorkersDone = collection.ChannelsMerge(worker1.Done(), worker2.Done())
	<-allWorkersDone

	fmt.Println("all workers finished")
}
```

### Collection Functions

This package provides several functions for working with collections:

*   `Aggregate[T, K any](source []T, aggregator func(K, T) K) K`
*   `Any[T any](source []T, predicate func(T) bool) bool`
*   `ChannelsMerge[T any](args ...<-chan T) <-chan T`
*   `ChannelsReadonly[T any](args ...chan T) []<-chan T`
*   `Contains[T comparable](source []T, item T) bool`
*   `Copy[T any](source []T) []T`
*   `Difference[T comparable](a []T, b []T) []T`
*   `Distinct[T comparable](source []T) []T`
*   `DistinctBy[T any](source []T, equals func(left T, right T) bool) []T`
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
*   `ChannelsMerge[T any](args ...<-chan T) <-chan T`
*   `ChannelsReadonly[T any](args ...chan T) []<-chan T`  
*   `MinOf[T constraints.Ordered](elements ...T) T`
*   `Min[T constraints.Ordered](l T, r T) T`
*   `Reverse[T any](source []T)`
*   `SliceToMap[K comparable, T any](source []T, keyFunc func(T) K) map[K]T`
*   `SortBy[T any](source []T, less func(l T, r T) bool)`
*   `Sort[T constraints.Ordered](source []T)`
*   `TransformBy[T, K any](source []T, transform func(T) K) []K`
*   `TransformManyBy[T, K any](source []T, transform func(T) []K) []K`
*   `TryTransformBy[T, K any](source []T, transform func(T) (K, error)) ([]K, error)`


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

