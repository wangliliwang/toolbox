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

type debounceState int

const (
	debounceStateNotTiming debounceState = 1
	debounceStateTiming    debounceState = 2
	debounceStateCanceled  debounceState = 3
)

// debounce implement debounce function.
// FIXME(@wangli) There're some problems in change states.
type debounce struct {
	// inner
	callCh   chan struct{}
	cancelCh chan struct{}
	flushCh  chan struct{}
	timer    *time.Timer
	mu       sync.RWMutex
	state    debounceState

	// outer
	f     func()
	delay time.Duration
}

func newDebounce(function func(), delay time.Duration) *debounce {
	timer := time.NewTimer(delay)
	timer.Stop()
	dbc := &debounce{
		callCh:   make(chan struct{}),
		cancelCh: make(chan struct{}),
		flushCh:  make(chan struct{}),
		timer:    timer,
		mu:       sync.RWMutex{},
		f:        function,
		delay:    delay,
	}
	// state machine
	go func() {
		defer dbc.closeResources()
		for {
			select {
			case <-dbc.callCh:
				dbc.onCall()
			case <-dbc.flushCh:
				dbc.onFlush()
			case <-dbc.timer.C:
				dbc.onTimer()
			case <-dbc.cancelCh:
				dbc.onCancel()
				return
			}
		}
	}()
	return dbc
}

func (d *debounce) closeResources() {
	close(d.callCh)
	close(d.cancelCh)
	close(d.flushCh)
	d.timer.Stop()
}

func (d *debounce) onCancel() {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.state == debounceStateCanceled {
		panic("can't cancel canceled debounce")
	}
	d.state = debounceStateCanceled
	d.timer.Stop()
	fmt.Println("onCancel: ", d.state)
}

func (d *debounce) onFlush() {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.state == debounceStateCanceled {
		panic("can't flush canceled debounce")
	}
	d.state = debounceStateNotTiming
	d.timer.Stop()
	d.f()
}

func (d *debounce) onCall() {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.state == debounceStateCanceled {
		panic("can't call canceled debounce")
	}
	d.state = debounceStateTiming
	d.timer.Reset(d.delay)
}

func (d *debounce) onTimer() {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.state == debounceStateCanceled || d.state == debounceStateNotTiming {
		panic("timer not stopped correctly")
	}

	d.state = debounceStateNotTiming
	d.timer.Stop()
	d.f()
}

func (d *debounce) cancel() {
	if d.state == debounceStateCanceled {
		panic("can't cancel canceled debounce")
	}
	d.cancelCh <- struct{}{}
}

func (d *debounce) flush() {
	if d.state == debounceStateCanceled {
		panic("can't flush canceled debounce")
	}
	fmt.Println("flush: ", d.state)
	d.flushCh <- struct{}{}
}

func (d *debounce) call() {
	if d.state == debounceStateCanceled {
		panic("can't call canceled debounce")
	}
	d.callCh <- struct{}{}
}

// NewDebounce returns call, flush and cancel function in debounce.
func NewDebounce(function func(), delay time.Duration) (call, flush, cancel func()) {
	dbc := newDebounce(function, delay)
	return dbc.call, dbc.flush, dbc.cancel
}
