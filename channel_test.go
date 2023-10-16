package toolbox

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChannelDispatcher(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	stream := make(chan int)
	go func() {
		for i := 0; i < 4; i++ {
			stream <- i
		}
		close(stream)
	}()
	channels := ChannelDispatcher(stream, 3, 0, DispatchingStrategyRoundRobin[int])
	result0 := <-channels[0]
	is.Equal(0, result0)
	result1 := <-channels[1]
	is.Equal(1, result1)
	result2 := <-channels[2]
	is.Equal(2, result2)
	result3 := <-channels[0]
	is.Equal(3, result3)
}

func TestDispatchingStrategyRoundRobin(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	channels := make([]<-chan int, 5)
	for i := range channels {
		channels[i] = make(chan int, 1)
	}
	for i := 0; i < 100; i++ {
		result := DispatchingStrategyRoundRobin(0, uint64(i), channels)
		is.Equal(i%len(channels), result)
	}
}

func TestDispatchingStrategyRandom(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	channels := make([]<-chan int, 5)
	for i := range channels {
		channels[i] = make(chan int, 1)
	}
	for i := 0; i < 100; i++ {
		result := DispatchingStrategyRandom(0, uint64(i), channels)
		is.True(result < len(channels))
	}
}

func TestDispatchingStrategyWeightedRandom(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	channels := make([]<-chan int, 5)
	for i := range channels {
		channels[i] = make(chan int, 1)
	}
	weights := []int{1, 2, 3, 2, 4}
	strategy := DispatchingStrategyWeightedRandom[int](weights)
	for i := 0; i < 100; i++ {
		result := strategy(0, uint64(i), channels)
		is.True(result < len(channels))
	}
}

func TestDispatchingStrategyFirst(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	channels := make([]chan int, 5)
	for i := range channels {
		channels[i] = make(chan int, 1)
	}
	channels[0] <- 0
	for i := 0; i < 100; i++ {
		result := DispatchingStrategyFirst(0, uint64(i), channelsToReadonly(channels))
		is.Equal(1, result)
	}
}

func TestDispatchingStrategyLeast(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	channels := make([]chan int, 5)
	for i := range channels {
		channels[i] = make(chan int, i)
		for j := 0; j < i; j++ {
			channels[i] <- j
		}
	}
	for i := 0; i < 100; i++ {
		result := DispatchingStrategyLeast(0, uint64(i), channelsToReadonly(channels))
		is.Equal(0, result)
	}
}

func TestDispatchingStrategyMost(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	channels := make([]chan int, 5)
	for i := range channels {
		channels[i] = make(chan int, 10)
		for j := 0; j < i; j++ {
			channels[i] <- j
		}
	}
	for i := 0; i < 100; i++ {
		result := DispatchingStrategyMost(0, uint64(i), channelsToReadonly(channels))
		is.Equal(4, result)
	}
}
