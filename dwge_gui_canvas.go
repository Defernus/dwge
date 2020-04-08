package dwge

import (
	"fmt"
)

type Canvas struct {
	x, y int
	w, h int
	parent GuiElement
}

func (c *Canvas) GetParent() GuiElement {
	return c.parent
}

func NewCanvas(x, y, w, h int, parent GuiContainer) (*Canvas, error) {
	c := &Canvas {
		x: x,
		y: y,
		w: w,
		h: h,
	}

	if _, err := parent.AddElement(c); err != nil {
		return nil, err
	}

	return c, nil
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
	return c.x+px, c.y+py
}
func (c *Canvas) GetSize() (int, int) {
	return c.w, c.h
}

func (c *Canvas) draw() error {
	//TODO make real draw here
	return nil
}

func (c *Canvas) Clear() {
	c.DrawRect(0, 0, c.w, c.h)
}

func (c *Canvas) DrawRect(x, y, w, h int) {
	//TODO move real draw calls to draw()

	px, py := c.GetAbsolutePosition()
	drawRect(px + x, py + y, w, h)
}

func (c *Canvas) DrawText(text string, x, y int) {
	//TODO move real draw calls to draw()

	px, py := c.GetAbsolutePosition()
	drawText(text, px + x, py + y)
}
