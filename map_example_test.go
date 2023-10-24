package toolbox

import (
	"fmt"
	"strings"
)

func ExampleOmitBy() {
	in := map[string]int{
		"apple":  10,
		"orange": 12,
		"banana": 15,
	}
	result := OmitBy(in, func(key string, value int) bool {
		return strings.HasPrefix(key, "or")
	})
	fmt.Printf("%+v %+v", result["apple"], result["banana"])
	// Output:
	// 10 15
}
