package toolbox

import (
	"math/rand"
	"sync"
	"time"
)

type DispatchingStrategy[T any] func(msg T, index uint64, channels []<-chan T) int

// ChannelDispatcher distributes messages from input channels into N child channels.
// Close events are propagated to children.
// Underlying channels can have a fixed buffer capacity or be unbuffered when cap is 0.
func ChannelDispatcher[T any](stream <-chan T, count int, channelBufferCap int, strategy DispatchingStrategy[T]) []<-chan T {
	channels := createChannels[T](count, channelBufferCap)
	roChildren := channelsToReadonly(channels)
	go func() {
		defer closeChannels(channels)
		var index uint64
		for {
			msg, ok := <-stream
			if !ok {
				break
			}
			destination := strategy(msg, index, roChildren) % count
			channels[destination] <- msg
			index++
		}
	}()
	return roChildren
}

func createChannels[T any](count, channelBufferCap int) []chan T {
	result := make([]chan T, 0, count)
	for i := 0; i < count; i++ {
		result = append(result, make(chan T, channelBufferCap))
	}
	return result
}

func channelsToReadonly[T any](channels []chan T) []<-chan T {
	result := make([]<-chan T, 0)
	for _, c := range channels {
		result = append(result, c)
	}
	return result
}

func closeChannels[T any](channels []chan T) {
	for _, c := range channels {
		close(c)
	}
}

func channelIsNotFull[T any](ch <-chan T) bool {
	return cap(ch) == 0 || len(ch) < cap(ch)
}

// DispatchingStrategyRoundRobin distributes messages in a rotating sequential manner.
// If the channel capacity is exceeded, the next channel will be selected and so on.
func DispatchingStrategyRoundRobin[T any](msg T, index uint64, channels []<-chan T) int {
	l := len(channels)
	for {
		result := int(index % uint64(l))
		if channelIsNotFull(channels[result]) {
			return result
		}
		time.Sleep(10 * time.Millisecond) // prevent CPU from burning 🔥
		index++
	}
}

// DispatchingStrategyRandom distributes messages in a random manner.
// If the channel capacity is exceeded, another random channel will be selected and so on.
func DispatchingStrategyRandom[T any](msg T, index uint64, channels []<-chan T) int {
	l := len(channels)
	for {
		result := rand.Intn(l)
		if channelIsNotFull(channels[result]) {
			return result
		}
		time.Sleep(10 * time.Millisecond) // prevent CPU from burning 🔥
	}
}

// DispatchingStrategyWeightedRandom distributes messages in a weighted manner.
// If the channel capacity is exceeded, another random channel will be selected and so on.
func DispatchingStrategyWeightedRandom[T any](weights []int) DispatchingStrategy[T] {
	// todo(wangli) optimize here.
	seq := make([]int, 0)
	for i := 0; i < len(weights); i++ {
		for j := 0; j < weights[i]; j++ {
			seq = append(seq, i)
		}
	}

	return func(msg T, index uint64, channels []<-chan T) int {
		l := len(seq)
		for {
			result := seq[rand.Intn(l)]
			if channelIsNotFull(channels[result]) {
				return result
			}
			time.Sleep(10 * time.Millisecond) // prevent CPU from burning 🔥
		}
	}
}

// DispatchingStrategyFirst distributes messages in a first-non-full manner.
// If the channel capacity is exceeded, the second channel will be selected and so on.
func DispatchingStrategyFirst[T any](msg T, index uint64, channels []<-chan T) int {
	for {
		for i, ch := range channels {
			if channelIsNotFull(ch) {
				return i
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}

// DispatchingStrategyLeast distributes messages in the emptiest channel.
func DispatchingStrategyLeast[T any](msg T, index uint64, channels []<-chan T) int {
	seq := Range(len(channels))
	return FindBy(seq, func(a, b int) bool {
		return len(channels[a]) < len(channels[b])
	})
}

// DispatchingStrategyMost distributes messages in the fullest channel.
// If the channel capacity is exceeded, the next channel will be selected and so on.
func DispatchingStrategyMost[T any](msg T, index uint64, channels []<-chan T) int {
	seq := Range(len(channels))
	return FindBy(seq, func(a, b int) bool {
		return (len(channels[a]) > len(channels[b])) && channelIsNotFull(channels[a])
	})
}

// SliceToChannel returns a read-only channel of collection items.
func SliceToChannel[T any](collection []T, bufferSize int) <-chan T {
	ch := make(chan T, bufferSize)
	go func() {
		for _, item := range collection {
			ch <- item
		}
		close(ch)
	}()
	return ch
}

// ChannelToSlice returns a slice built from channel items.
// Blocks until channel closes.
func ChannelToSlice[T any](ch <-chan T) []T {
	collection := make([]T, 0)
	for item := range ch {
		collection = append(collection, item)
	}
	return collection
}

// Generator implements the generator design pattern.
// Refer: https://github.com/tmrts/go-patterns/blob/master/concurrency/generator.md
func Generator[T any](bufferSize int, generator func(yield func(T))) <-chan T {
	ch := make(chan T, bufferSize)
	go func() {
		// WARNING: infinite loop
		generator(func(t T) {
			ch <- t
		})
		close(ch)
	}()
	return ch
}

// Buffer creates a slice of n elements from a channel. Returns the slice and slice length.
func Buffer[T any](ch <-chan T, size int) (collection []T, length int, readTime time.Duration, closed bool) {
	collection = make([]T, 0, size)
	now := time.Now()
	closed = true
	for index := 0; index < size; index++ {
		value, ok := <-ch
		if !ok {
			closed = false
			break
		}
		collection = append(collection, value)
	}
	length = len(collection)
	readTime = time.Since(now)
	return
}

// BufferWithTimeout
func BufferWithTimeout[T any](ch <-chan T, size int, timeout time.Duration) (collection []T, length int, readTime time.Duration, closed bool) {
	expire := time.NewTimer(timeout)
	defer expire.Stop()

	collection = make([]T, 0, size)
	now := time.Now()
	closed = true
	for index := 0; index < size; index++ {
		select {
		case value, ok := <-ch:
			if !ok {
				closed = false
				break
			}
			collection = append(collection, value)
		case <-expire.C:
			break
		}
	}
	length = len(collection)
	readTime = time.Since(now)
	return
}

// FanIn collects items from multiple channels into a single buffered channel.
// Output messages has no priority. When all upstream channels reach EOF, downstream channel closes.
func FanIn[T any](channelBufferCap int, upstreams ...<-chan T) <-chan T {
	out := make(chan T, channelBufferCap)
	wg := sync.WaitGroup{}
	wg.Add(len(upstreams))
	for _, upstream := range upstreams {
		go func(ch <-chan T) {
			for item := range ch {
				out <- item
			}
			wg.Done()
		}(upstream)
	}

	go func() {
		wg.Wait()
		close(out) // TODO(@wangli) reader has problems?
	}()

	return out
}

// FanOut broadcasts all upstream messages to multiple downstream channels.
// When upstream channel EOF, downstream channels close.
// If any downstream channel is full, broadcasting is paused.
func FanOut[T any](count, channelBufferCap int, upstream <-chan T) []<-chan T {
	downstreamChannels := createChannels[T](count, channelBufferCap)

	go func() {
		for item := range upstream {
			for _, ch := range downstreamChannels {
				ch <- item
			}
		}
		closeChannels(downstreamChannels)
	}()

	return channelsToReadonly(downstreamChannels)
}
