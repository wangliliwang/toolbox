package toolbox

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCamelCase(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	raw1 := "hello, world! It's wonderful."
	is.Equal("helloWorldItSWonderful", CamelCase(raw1))
}
