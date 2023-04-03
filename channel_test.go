package collection_test

import (
	"testing"

	"github.com/sergeydobrodey/collection"
)

func TestChannelsMerge(t *testing.T) {
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

	const expected = 10

	if sum != expected {
		t.Errorf("ChannelsMerge returned %v, expected %v", sum, expected)
	}
}
