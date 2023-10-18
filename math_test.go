package toolbox

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRange(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Range[int](0)
	is.Equal([]int{}, result1)

	result2 := Range[int](3)
	is.Equal([]int{0, 1, 2}, result2)
}
