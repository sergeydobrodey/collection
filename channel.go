package collection

import "sync"

// ChannelsReadonly transforms input N channels to receive only channels
func ChannelsReadonly[T any](args ...chan T) []<-chan T {
	return TransformBy(args, func(v chan T) <-chan T {
		return v
	})
}

// ChannelsMerge merge input from N channels to 1 receive only channel
func ChannelsMerge[T any](args ...<-chan T) <-chan T {
	result := make(chan T)

	wg := sync.WaitGroup{}
	wg.Add(len(args))

	go func() {
		wg.Wait()
		close(result)
	}()

	for _, c := range args {
		go func(c <-chan T) {
			for v := range c {
				result <- v
			}
			wg.Done()
		}(c)
	}

	return result
}
