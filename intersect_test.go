package toolbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
