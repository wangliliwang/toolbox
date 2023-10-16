package toolbox

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	min := FindBy([]int{1, 2, 3, 4, 5}, func(a, b int) bool {
		return a < b
	})
	is.Equal(1, min)

	max := FindBy([]int{1, 2, 3, 4, 5}, func(a, b int) bool {
		return a > b
	})
	is.Equal(5, max)

	zero := FindBy([]int{}, func(a, b int) bool {
		return a > b
	})
	is.Equal(0, zero)
}
