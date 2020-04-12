package dwge

import (
	"fmt"
	"syscall/js"
)

var (
	document js.Value
	alert    js.Value
)

//init helper and Get() all essential js elements
func initWH(width, height int) {
	document = js.Global().Get("document")
	alert = js.Global().Get("alert")

	screen = &Image{}

	screen.canvas = document.Call("getElementById", "dwge_canvas")

	screen.canvas.Set("width", width)
	screen.canvas.Set("height", height)

	screen.w = width
	screen.h = height

	screen.ctx = screen.canvas.Call("getContext", "2d")
	screen.ctx.Set("imageSmoothingEnabled", false)
}

func startLoop(f func() error) {
	var js_f js.Func
	js_f = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if err := f(); err != nil {
			//TODO handle loop error
			fmt.Printf("loop error: %s\n", err)
		}
		js.Global().Call("requestAnimationFrame", js_f)
		return nil
	})
	js.Global().Call("requestAnimationFrame", js_f)
}
