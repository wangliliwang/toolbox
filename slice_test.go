package toolbox

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO(@wangli) edge case

func TestFilter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Filter([]int{1, 2, 3, 4, 5}, func(item int, index int) bool {
		return item%2 == 0
	})
	is.Equal([]int{2, 4}, result1)

	result2 := Filter([]int{}, func(item int, index int) bool {
		return item%2 == 0
	})
	is.Equal([]int{}, result2)
}

func TestContains(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Contains([]int{1, 2, 3, 4, 5}, 5)
	result2 := Contains([]int{1, 2, 3, 4, 5}, 6)

	is.True(result1)
	is.False(result2)
}

func TestContainsBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	type a struct {
		A int
		B string
	}

	data := []a{{A: 1, B: "a"}, {A: 2, B: "B"}, {A: 3, B: "C"}}
	result1 := ContainsBy(data, func(item a) bool { return item.A == 1 && item.B == "a" })
	result2 := ContainsBy(data, func(item a) bool { return item.A == 1 && item.B == "B" })

	is.True(result1)
	is.False(result2)
}

func TestEvery(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Every([]int{1, 2, 3, 4, 5}, []int{2})
	result2 := Every([]int{1, 2, 3, 4, 5}, []int{2, 3, 4})
	result3 := Every([]int{1, 2, 3, 4, 5}, []int{23, 3, 4})

	is.True(result1)
	is.True(result2)
	is.False(result3)
}

func TestEveryBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := EveryBy([]int{1, 2, 3, 4, 5}, func(item int) bool { return item <= 6 })
	result2 := EveryBy([]int{1, 2, 3, 4, 5, 7}, func(item int) bool { return item <= 6 })

	is.True(result1)
	is.False(result2)
}

func TestSome(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Some([]int{1, 2, 3, 4, 5}, []int{2})
	result2 := Some([]int{1, 2, 3, 4, 5}, []int{2, 3, 6})
	result3 := Some([]int{1, 2, 3, 4, 5}, []int{6, 7, 8})

	is.True(result1)
	is.True(result2)
	is.False(result3)
}

func TestSomeBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := SomeBy([]int{1, 2, 3, 4, 5}, func(item int) bool { return item <= 3 })
	result2 := SomeBy([]int{1, 2, 3, 4, 5, 7}, func(item int) bool { return item <= -1 })

	is.True(result1)
	is.False(result2)
}

func TestNone(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := None([]int{1, 2, 3, 4, 5}, []int{2})
	result2 := None([]int{1, 2, 3, 4, 5}, []int{2, 3, 6})
	result3 := None([]int{1, 2, 3, 4, 5}, []int{6, 7, 8})

	is.False(result1)
	is.False(result2)
	is.True(result3)
}

func TestNoneBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := NoneBy([]int{1, 2, 3, 4, 5}, func(item int) bool { return item <= 3 })
	result2 := NoneBy([]int{1, 2, 3, 4, 5, 7}, func(item int) bool { return item <= -1 })

	is.False(result1)
	is.True(result2)
}
