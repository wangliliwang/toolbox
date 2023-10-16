package toolbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestIntersect(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Intersect([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 6, 7})
	result2 := Intersect([]int{1, 2, 3, 4, 5}, []int{6, 7, 8})

	is.Equal([]int{3, 4, 5}, result1)
	is.Equal([]int{}, result2)
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Subtract([]int{}, []int{1, 2, 3})
	result2 := Subtract([]int{1}, []int{1, 2, 3})
	result3 := Subtract([]int{1, 2, 3, 4}, []int{1, 2, 3})

	is.Equal([]int{}, result1)
	is.Equal([]int{}, result2)
	is.Equal([]int{4}, result3)
}

func TestUnion(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Union([]int{1, 1, 2, 3}, []int{1, 2, 3, 4})
	result2 := Union([]int{1, 1, 2, 3}, []int{})

	is.Equal([]int{1, 2, 3, 4}, result1)
	is.Equal([]int{1, 2, 3}, result2)
}

func TestWithout(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Without([]int{1, 1, 2, 3}, 1, 2)
	result2 := Without([]int{1, 1, 2, 3})

	is.Equal([]int{3}, result1)
	is.Equal([]int{1, 1, 2, 3}, result2)
}

func TestWithoutEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := WithoutEmpty([]int{1, 1, 2, 3})
	result2 := WithoutEmpty([]int{0, 1, 1, 2, 3})

	is.Equal([]int{1, 1, 2, 3}, result1)
	is.Equal([]int{1, 1, 2, 3}, result2)
}
