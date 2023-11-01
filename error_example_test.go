package toolbox

import (
	"errors"
	"fmt"
)

func wrapPanicIf(srcErr error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = fmt.Errorf("%w", e)
			} else {
				err = fmt.Errorf("unknown panic")
			}
		}
	}()
	PanicIf(srcErr)
	return nil
}

func ExamplePanicIf() {
	var err error
	fmt.Println(wrapPanicIf(err))
	err = errors.New("example err")
	fmt.Println(wrapPanicIf(err))
	// Output:
	// <nil>
	// example err
}
