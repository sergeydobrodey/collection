package collection_test

import (
	"testing"

	"github.com/sergeydobrodey/collection"
)

func TestChannelsMerge(t *testing.T) {
	const expected = 10

	var sources []chan int

	for i := 0; i < 5; i++ {
		sources = append(sources, make(chan int, 1))
		sources[i] <- i
		close(sources[i])
	}

	var merged = collection.ChannelsMerge(collection.ChannelsReadonly(sources...)...)

	var sum int
	for v := range merged {
		sum += v
	}

	if sum != expected {
		t.Errorf("ChannelsMerge = %v; want %v", sum, expected)
	}
}
