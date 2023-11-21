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

type debounceAction int

const (
	debounceActionCall   debounceAction = 1
	debounceActionFlush  debounceAction = 2
	debounceActionCancel debounceAction = 3
)

// debounce implement debounce function.
// FIXME(@wangli) There're some problems in change states.
type debounce struct {
	// inner
	actionCh chan debounceAction
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
		actionCh: make(chan debounceAction),
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
			case v, ok := <-dbc.actionCh:
				if !ok {
					return
				}
				switch v {
				case debounceActionCall:
					dbc.onCall()
				case debounceActionFlush:
					dbc.onFlush()
				case debounceActionCancel:
					dbc.onCancel()
					return
				}
			case <-dbc.timer.C:
				dbc.onTimer()
			}
		}
	}()
	return dbc
}

func (d *debounce) closeResources() {
	close(d.actionCh)
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
	d.actionCh <- debounceActionCancel
}

func (d *debounce) flush() {
	d.actionCh <- debounceActionFlush
}

func (d *debounce) call() {
	d.actionCh <- debounceActionCall
}

// NewDebounce returns call, flush and cancel function in debounce.
func NewDebounce(function func(), delay time.Duration) (call, flush, cancel func()) {
	dbc := newDebounce(function, delay)
	return dbc.call, dbc.flush, dbc.cancel
}

// Memoize creates a function that memoize the result of func.
func Memoize[T any](function func(key string) T) func(key string) T {
	cache := make(map[string]T)
	return func(key string) T {
		if v, ok := cache[key]; ok {
			return v
		} else {
			calculatedV := function(key)
			cache[key] = calculatedV
			return calculatedV
		}
	}
}
