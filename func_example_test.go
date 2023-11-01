package toolbox

import (
	"errors"
	"fmt"
)

func ExampleSafelyRun() {
	err := SafelyRun(func() {
		panic(errors.New("example err"))
	})
	fmt.Println(err)
}

func ExampleAfter() {
	result := After(3, func() {
		fmt.Println("info")
	})
	result()
	result()
	result()
	result()
	// Output:
	// info
	// info
}
