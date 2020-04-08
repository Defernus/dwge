package dwge

import (
	"syscall/js"
)

func SetKeyPressEvent(f func(key string)) {
	js_f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f(args[0].Get("key").String())
		return nil
	})

	document.Call("addEventListener", "keypress", js_f)
}
