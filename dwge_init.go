package dwge

import (
	"time"
)

var (
	screen *Image
)

func getMainLoop(loop func(GuiElement) error, mwin *mainWindow) func() error {
	last_time := time.Now()
	return func() error {
		now := time.Now()
		delta_time = float64((now.UnixNano()-last_time.UnixNano())/1000000) / 1000.
		last_time = now

		if err := loop(mwin); err != nil {
			return err
		}

		screen.SetFillColor(0, 0, 0)
		screen.Clear()
		mwin.draw()
		return nil
	}
}

//Init starts main loop
func Init(create func(GuiElement) error, loop func(GuiElement) error, width, height int) error {
	done := make(chan struct{}, 0)

	initWH(width, height)
	initEvents()

	mwin := newMainWindow()

	if err := create(mwin); err != nil {
		return err
	}

	startLoop(getMainLoop(loop, mwin))

	<-done
	clearEventsListeners()
	return nil
}
