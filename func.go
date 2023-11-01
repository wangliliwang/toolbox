package toolbox

import (
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

// SafelyRun try to catch panic during f-runtime, and transform it into error.
func SafelyRun(function func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = fmt.Errorf("%w\n%s", e, debug.Stack())
			} else {
				err = fmt.Errorf("unknown error\n%s", debug.Stack())
			}
		}
	}()
	function()
	return nil
}

// After creates a function that invokes input `function` once it's called n or more times.
func After(n int, function func()) func() {
	var count int
	return func() {
		count++
		if count >= n {
			function()
		}
	}
}

type debounce struct {
	// inner
	callCh   chan struct{}
	cancelCh chan struct{}
	flushCh  chan struct{}
	timer    *time.Timer
	mu       sync.Mutex

	// outer
	f     func()
	delay time.Duration
}

func newDebounce(function func(), delay time.Duration) *debounce {
	dbc := &debounce{
		callCh:   make(chan struct{}),
		cancelCh: make(chan struct{}),
		flushCh:  make(chan struct{}),
		mu:       sync.Mutex{},
		f:        function,
	}
	callAndClose := func() {
		// 关闭计时
		if dbc.timer != nil {
			dbc.timer.Stop()
		}
		// 立刻调用
		dbc.f()
		dbc.timer = time.NewTimer(dbc.delay)
	}
	go func() {
		defer func() {
			close(callCh)
			close(cancelCh)
			if timer != nil {
				timer.Stop()
			}
		}()
		for {
			select {
			case <-callCh:
				// 重新计时
				if timer == nil {
					timer = time.NewTimer(delay)
				} else {
					timer.Reset(delay)
				}
			case <-flushCh:
				callAndClose()
			case <-timer.C:
				callAndClose()
			case <-cancelCh:
				// 取消调用
				return
			}
		}
	}()
}

func (d *debounce) cancel() {
	d.cancelCh <- struct{}{}
}

func (d *debounce) flush() {
	d.flushCh <- struct{}{}
}

func (d *debounce) call() {
	d.callCh <- struct{}{}
}
