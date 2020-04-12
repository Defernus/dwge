package dwge

import (
	"fmt"
	"syscall/js"
)

type Image struct {
	canvas, ctx js.Value
	w, h        int
}

func LoadImage(url string) (*Image, error) {
	var img *Image

	js_img := js.Global().Get("Image").New()

	done := make(chan error)

	js_load := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w, h := js_img.Get("width").Int(), js_img.Get("height").Int()

		img = NewImage(w, h)

		img.ctx.Call("drawImage", js_img, 0, 0)

		done <- nil
		return nil
	})

	js_error := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		done <- fmt.Errorf("failed to load image from %s\n", url)
		return nil
	})

	js_img.Call("addEventListener", "load", js_load)
	js_img.Call("addEventListener", "error", js_error)
	js_img.Set("src", url)

	err := <-done
	return img, err
}

func NewImage(w, h int) *Image {
	img := &Image{}
	img.w = w
	img.h = h
	img.canvas = document.Call("createElement", "canvas")
	img.canvas.Set("width", w)
	img.canvas.Set("height", h)
	img.ctx = img.canvas.Call("getContext", "2d")
	img.ctx.Set("imageSmoothingEnabled", false)
	return img
}

func (img *Image) Clear() {
	img.DrawRect(0, 0, img.w, img.h)
}

func (img *Image) SetFillColor(r, g, b float64) {
	img.ctx.Set("fillStyle", fmt.Sprintf("rgb(%v, %v, %v)", int(r*255), int(g*255), int(b*255)))
}

func (img *Image) DrawRect(x, y, w, h int) {
	img.ctx.Call("fillRect", x, img.h-y-h, w, h)
}

func (img *Image) SetFontSize(px uint8) {
	img.ctx.Set("font", fmt.Sprintf("%vpx Verdana", px))
}

func (img *Image) DrawText(text string, x, y int) {
	img.ctx.Call("fillText", text, x, img.h-y)
}

func (img *Image) DrawImage(image *Image) {
	img.ctx.Call("drawImage", image.canvas, 0, img.h-image.h)
}

func (img *Image) DrawImageAt(image *Image, dx, dy int) {
	img.ctx.Call("drawImage", image.canvas, dx, img.h-dy-image.h)
}

func (img *Image) DrawScaledImageAt(image *Image, dx, dy, dw, dh int) {
	img.ctx.Call("drawImage", image.canvas, dx, img.h-dy-dh, dw, dh)
}

func (img *Image) DrawScaledPartOfImageAt(image *Image, sx, sy, sw, sh, dx, dy, dw, dh int) {
	img.ctx.Call("drawImage", image.canvas, sx, image.h-sy-sh, sw, sh, dx, img.h-dy-dw, dw, dh)
}

func (img *Image) DrawRotatedScaledImageAt(image *Image, dx, dy, dw, dh int, cx, cy, angle float64) {
	dy = img.h - dy
	cy = 1 - cy

	img.ctx.Call("translate", float64(dx), float64(dy))
	img.ctx.Call("rotate", -angle)
	img.ctx.Call("translate", -cx*float64(dw), -cy*float64(dh))

	img.ctx.Call("drawImage", image.canvas, 0, 0, dw, dh)

	img.ctx.Call("resetTransform")
}

func (img *Image) DrawRotatedScaledPartOfImageAt(image *Image, sx, sy, sw, sh, dx, dy, dw, dh int, cx, cy, angle float64) {
	sy = image.h - sy - sh
	dy = img.h - dy
	cy = 1 - cy

	img.ctx.Call("translate", float64(dx), float64(dy))
	img.ctx.Call("rotate", -angle)
	img.ctx.Call("translate", -cx*float64(dw), -cy*float64(dh))

	img.ctx.Call("drawImage", image.canvas, sx, sy, sh, sw, 0, 0, dw, dh)

	img.ctx.Call("resetTransform")
}
