package toolbox

import (
	"math/rand"
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
