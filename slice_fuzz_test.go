package toolbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// See https://go.dev/security/fuzz/

func FuzzReverse(f *testing.F) {
	f.Add([]byte{1, 3, 4, 5, 6})
	f.Fuzz(func(t *testing.T, list []byte) {
		listCopyed := make([]byte, len(list))
		copy(listCopyed, list)
		Reverse(list)
		Reverse(list)
		assert.Equal(t, listCopyed, list)
	})
}
