package toolbox

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

var lengths = []int{10, 100, 1000}

func BenchmarkChunk(b *testing.B) {

	for _, n := range lengths {
		strs := genSliceString(n)
		b.Run(fmt.Sprintf("strings_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Chunk(strs, 5)
			}
		})
	}

	for _, n := range lengths {
		ints := genSliceInt(n)
		b.Run(fmt.Sprintf("ints_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Chunk(ints, 5)
			}
		})
	}
}

func genSliceString(n int) []string {
	result := make([]string, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, strconv.Itoa(rand.Intn(100_000)))
	}
	return result
}

func genSliceInt(n int) []int {
	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, rand.Intn(100_000))
	}
	return result
}
