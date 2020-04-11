package dwge

import (
	"fmt"
	"syscall/js"
)

type Image struct {
	canvas, ctx js.Value
	w, h        int
}

func NewImage(w, h int) *Image {
	img := &Image{}
	img.w = w
	img.h = h
	img.canvas = document.Call("createElement", "canvas")
	img.canvas.Set("width", w)
	img.canvas.Set("height", h)
	img.ctx = img.canvas.Call("getContext", "2d")
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
