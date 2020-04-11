package dwge

import (
	"syscall/js"
)

var (
	key_down_event    func(string)
	js_key_down_event js.Func

	key_up_event    func(string)
	js_key_up_event js.Func

	lmb_click_event    func(x, y int)
	js_lmb_click_event js.Func

	rmb_click_event    func(x, y int)
	js_rmb_click_event js.Func

	mouseover_event    func(x, y int)
	js_mouseover_event js.Func

	mouseout_event    func(x, y int)
	js_mouseout_event js.Func

	mousedown_event    func(button, x, y int)
	js_mousedown_event js.Func

	mouseup_event    func(button, x, y int)
	js_mouseup_event js.Func

	mousemove_event    func(x, y, dx, dy int)
	js_mousemove_event js.Func
)

func initEvents() {
	js_key_down_event = js.FuncOf(keyDownEvent)
	document.Call("addEventListener", "keydown", js_key_down_event)

	js_key_up_event = js.FuncOf(keyDownEvent)
	document.Call("addEventListener", "keyup", js_key_down_event)

	js_rmb_click_event = js.FuncOf(rmbClickEvent)
	document.Call("addEventListener", "click", js_rmb_click_event)

	js_lmb_click_event = js.FuncOf(lmbClickEvent)
	document.Call("addEventListener", "contextmenu", js_lmb_click_event)

	js_mouseover_event = js.FuncOf(mouseoverEvent)
	document.Call("addEventListener", " mouseover", js_mouseover_event)

	js_mouseout_event = js.FuncOf(mouseoutEvent)
	document.Call("addEventListener", " mouseout", js_mouseout_event)

	js_mousedown_event = js.FuncOf(mousedownEvent)
	document.Call("addEventListener", " mousedown", js_mousedown_event)

	js_mouseup_event = js.FuncOf(mouseupEvent)
	document.Call("addEventListener", " mouseup", js_mouseup_event)

	js_mousemove_event = js.FuncOf(mousemoveEvent)
	document.Call("addEventListener", " mousemove", js_mousemove_event)
}

func keyDownEvent(this js.Value, args []js.Value) interface{} {
	if key_down_event != nil {
		key_down_event(args[0].Get("key").String())
	}
	return nil
}

func SetKeyDownEvent(f func(key string)) {
	key_down_event = f
}

func keyUpEvent(this js.Value, args []js.Value) interface{} {
	if key_up_event != nil {
		key_up_event(args[0].Get("key").String())
	}
	return nil
}

func SetKeyUpEvent(f func(key string)) {
	key_up_event = f
}

func lmbClickEvent(this js.Value, args []js.Value) interface{} {
	if lmb_click_event != nil {
		lmb_click_event(args[0].Get("offsetX").Int(), args[0].Get("offsetY").Int())
	}
	return nil
}

func rmbClickEvent(this js.Value, args []js.Value) interface{} {
	if rmb_click_event != nil {
		rmb_click_event(args[0].Get("offsetX").Int(), args[0].Get("offsetY").Int())
	}
	return nil
}

func mouseoverEvent(this js.Value, args []js.Value) interface{} {
	if mouseover_event != nil {
		mouseover_event(args[0].Get("offsetX").Int(), args[0].Get("offsetY").Int())
	}
	return nil
}

func mouseoutEvent(this js.Value, args []js.Value) interface{} {
	if mouseout_event != nil {
		mouseout_event(args[0].Get("offsetX").Int(), args[0].Get("offsetY").Int())
	}
	return nil
}

func mousedownEvent(this js.Value, args []js.Value) interface{} {
	if mousedown_event != nil {
		mousedown_event(args[0].Get("button").Int(), args[0].Get("offsetX").Int(), args[0].Get("offsetY").Int())
	}
	return nil
}

func mouseupEvent(this js.Value, args []js.Value) interface{} {
	if mouseup_event != nil {
		mouseup_event(args[0].Get("button").Int(), args[0].Get("offsetX").Int(), args[0].Get("offsetY").Int())
	}
	return nil
}

func mousemoveEvent(this js.Value, args []js.Value) interface{} {
	if mousemove_event != nil {
		mousemove_event(args[0].Get("offsetX").Int(), args[0].Get("offsetY").Int(), args[0].Get("movementX").Int(), args[0].Get("movementY").Int())
	}
	return nil
}

func clearEventsListeners() {
	js_key_down_event.Release()
	js_key_up_event.Release()
}
