package dwge

import (
	"syscall/js"
)

var (
	document js.Value
	alert    js.Value
	canvas   js.Value
	ctx      js.Value
)

//init helper and Get() all essential js elements
func initWH(width, height int) {
	document = js.Global().Get("document")
	alert = js.Global().Get("alert")

	canvas = document.Call("getElementById", "dwge_canvas")

	canvas.Set("width", width)
	canvas.Set("height", height)

	ctx = canvas.Call("getContext", "2d")
}

func startLoop(f func() error) {
	var js_f js.Func
	js_f = js.FuncOf(func (this js.Value, args []js.Value) interface{}{
		if err := f(); err != nil {
			//TODO handle loop error
		}
		js.Global().Call("requestAnimationFrame", js_f)
		return nil
	})
	js.Global().Call("requestAnimationFrame", js_f)
}

func getCanvasSize() (int, int) {
	return canvas.Get("width").Int(), canvas.Get("height").Int()
}
