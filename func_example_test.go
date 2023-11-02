package toolbox

import (
	"errors"
	"fmt"
	"sync"
	"time"
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

func ExampleNewDebounce() {
	call, flush, _ := NewDebounce(func() {
		fmt.Println("exec")
	}, 300*time.Millisecond)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(100 * time.Millisecond)
			call()
		}
		flush()
		for i := 0; i < 3; i++ {
			time.Sleep(100 * time.Millisecond)
			call()
		}
		wg.Done()
	}()
	wg.Wait()
	time.Sleep(500 * time.Millisecond)
	// Output:
	// exec
	// exec
}
