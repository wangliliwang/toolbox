package toolbox

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

func TestGenerator(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Eventually(func() bool {
		ch := Generator[int](3, func(yield func(t int)) {
			for i := 0; i < 3; i++ {
				yield(i)
			}
		})
		index := 0
		for item := range ch {
			is.Equal(index, item)
			index++
		}
		return true
	}, time.Second, 10*time.Millisecond)
}

func TestBuffer(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Eventually(func() bool {
		ch := make(chan int)
		go func() {
			for i := 0; i < 8; i++ {
				ch <- i
			}
			close(ch)
		}()

		// not close
		collection1, length1, _, ok1 := Buffer(ch, 4)
		is.Equal([]int{0, 1, 2, 3}, collection1)
		is.Equal(4, length1)
		is.True(ok1)

		// closed
		collection2, length2, _, ok2 := Buffer(ch, 5)
		is.Equal([]int{4, 5, 6, 7}, collection2)
		is.Equal(4, length2)
		is.False(ok2)

		return true
	}, time.Second, 10*time.Millisecond)
}

func TestBufferWithTimeout(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Eventually(func() bool {
		ch := make(chan int)
		go func() {
			for i := 0; i < 9; i++ {
				if i != 0 && i%4 == 0 {
					time.Sleep(100 * time.Millisecond)
				}
				ch <- i
			}
			close(ch)
		}()

		// not close, not timeout
		collection1, length1, _, ok1 := BufferWithTimeout(ch, 4, 10*time.Millisecond)
		is.Equal([]int{0, 1, 2, 3}, collection1)
		is.Equal(4, length1)
		is.True(ok1)
		time.Sleep(100 * time.Millisecond)

		// not close, timeout
		collection2, length2, _, ok2 := BufferWithTimeout(ch, 5, 10*time.Millisecond)
		is.Equal([]int{4, 5, 6, 7}, collection2)
		is.Equal(4, length2)
		is.True(ok2)
		time.Sleep(100 * time.Millisecond)

		// close
		collection3, length3, _, ok3 := BufferWithTimeout(ch, 5, 10*time.Millisecond)
		is.Equal([]int{8}, collection3)
		is.Equal(1, length3)
		is.False(ok3)

		return true
	}, time.Second, 10*time.Millisecond)
}

func TestFanIn(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Eventually(func() bool {
		channels := createChannels[int](3, 0)
		for index, item := range channels {
			go func(value int, ch chan int) {
				ch <- value
				close(ch)
			}(index, item)
		}
		out := FanIn(0, channelsToReadonly(channels)...)
		collection := ChannelToSlice(out)
		is.Equal(3, len(collection))
		return true
	}, time.Second, 10*time.Millisecond)
}

func TestFanOut(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	is.Eventually(func() bool {
		upstream := make(chan int, 5)
		upstream <- 1
		upstream <- 2
		upstream <- 3
		close(upstream)
		downstreams := FanOut(3, 5, upstream)
		for _, downstream := range downstreams {
			index := 1
			for item := range downstream {
				is.Equal(index, item)
				index++
			}
		}
		return true
	}, time.Second, 10*time.Millisecond)
}

func TestSliceToChannel(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	is.Eventually(func() bool {
		collection := []int{1, 2, 3, 4, 5}
		ch := SliceToChannel(collection, 0)
		outCollection := ChannelToSlice(ch)
		is.Equal(collection, outCollection)
		return true
	}, time.Second, 10*time.Millisecond)
}
