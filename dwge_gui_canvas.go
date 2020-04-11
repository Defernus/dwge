package dwge

import (
	"fmt"
)

type Canvas struct {
	x, y   int
	w, h   int
	img    *Image
	parent GuiElement
}

func NewCanvas(x, y, w, h int, parent GuiContainer) (*Canvas, error) {
	c := &Canvas{
		x:   x,
		y:   y,
		w:   w,
		h:   h,
		img: NewImage(w, h),
	}

	if _, err := parent.AddElement(c); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Canvas) GetParent() GuiElement {
	return c.parent
}

func (c *Canvas) setParent(parent GuiElement) error {
	if c.parent != nil {
		return fmt.Errorf("element alredy has an parent")
	}
	c.parent = parent
	return nil
}

func (c *Canvas) clearParent() {
	c.parent = nil
}

func (c *Canvas) GetRelativePosition() (int, int) {
	return c.x, c.y
}

func (c *Canvas) GetAbsolutePosition() (int, int) {
	if c.parent == nil {
		return c.x, c.y
	}
	px, py := c.parent.GetAbsolutePosition()
	return c.x + px, c.y + py
}
func (c *Canvas) GetSize() (int, int) {
	return c.w, c.h
}

func (c *Canvas) draw() error {
	screen.ctx.Call("drawImage", c.img.canvas, 0, 0)
	return nil
}

func (c *Canvas) GetImage() (*Image, error) {
	if c.img == nil {
		return nil, fmt.Errorf("image is not set, try to use NewCanvas")
	}
	return c.img, nil
}
