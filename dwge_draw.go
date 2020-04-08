package dwge

import (
	"fmt"
)

func SetFillColor(r, g, b float64) {
	ctx.Set("fillStyle", fmt.Sprintf("rgb(%v, %v, %v)", int(r*255), int(g*255), int(b*255)))
}

func drawRect(x, y, w, h int) {
	ctx.Call("fillRect", x, canvas.Get("height").Int() - y - h, w, h)
}

func SetFontSize(px uint8) {
	ctx.Set("font", fmt.Sprintf("%vpx Verdana", px))
}

func drawText(text string, x, y int) {
	ctx.Call("fillText", text, x, canvas.Get("height").Int() - y)
}
