package dwge

import (
	"sync"
	"time"
)

var (
	main_mutex          sync.Mutex
	screen              *Image
	is_cocnurrency_safe = true
)

func getMainLoop(loop func(*MainWindow) error, mwin *MainWindow) func() error {
	last_time := time.Now()
	return func() error {
		if is_cocnurrency_safe {
			main_mutex.Lock()
		}
		now := time.Now()
		delta_time = float64((now.UnixNano()-last_time.UnixNano())/1000000) / 1000.
		last_time = now

		if err := loop(mwin); err != nil {
			return err
		}

		screen.SetFillColor(0, 0, 0)
		screen.Clear()
		mwin.draw()

		if is_cocnurrency_safe {
			defer main_mutex.Unlock()
		}
		return nil
	}
}

func GetIsConcurrencySafe() bool {
	return is_cocnurrency_safe
}

func IsConcurencySafe(val bool) {
	is_cocnurrency_safe = true
}

//Init starts main loop
func Init(create func(*MainWindow) error, loop func(*MainWindow) error, width, height int) error {
	done := make(chan struct{}, 0)

	initWH(width, height)

	mwin := newMainWindow()
	mousedown_event = func(button, x, y int) { mwin.mouseButtonDownEvent(x, y, button) }
	mouseup_event = func(button, x, y int) { mwin.mouseButtonUpEvent(x, y, button) }
	initEvents()

	if err := create(mwin); err != nil {
		return err
	}

	startLoop(getMainLoop(loop, mwin))

	<-done
	clearEventsListeners()
	return nil
}
