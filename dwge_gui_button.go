package dwge

import (
	"fmt"
)

type Button struct {
	x, y        int
	w, h        int
	fnt_size    int
	color       string
	text        string
	event       func()
	img         *Image
	pressed_img *Image
	is_pressed  bool
	parent      GuiElement
}

func NewButton(img *Image, pressed_img *Image, event func(), x, y, w, h int, parent GuiContainer) (*Button, error) {
	button := &Button{
		x:           x,
		y:           y,
		w:           w,
		h:           h,
		event:       event,
		img:         img,
		pressed_img: pressed_img,
	}

	if err := parent.AddElement(button); err != nil {
		return nil, err
	}

	return button, nil
}

func (button *Button) mouseButtonDownEvent(x, y, b int) bool {
	if b != 0 {
		return false
	}

	ax, ay := button.GetAbsolutePosition()
	x -= ax
	y -= ay
	if x >= 0 && x < button.w && y >= 0 && y < button.h {
		button.is_pressed = true
		if button.event != nil {
			button.event()
		}
		return true
	}
	return false
}

func (button *Button) mouseButtonUpEvent(x, y, b int) {
	if b == MOUSE_BUTTON_LEFT {
		button.is_pressed = false
	}
}

func (button *Button) GetParent() GuiElement {
	return button.parent
}

func (button *Button) setParent(parent GuiElement) error {
	if button.parent != nil {
		return fmt.Errorf("element alredy has an parent")
	}
	button.parent = parent
	return nil
}

func (button *Button) clearParent() {
	button.parent = nil
}

func (button *Button) GetRelativePosition() (int, int) {
	return button.x, button.y
}

func (button *Button) GetAbsolutePosition() (int, int) {
	px, py := button.parent.GetAbsolutePosition()
	return px + button.x, py + button.y
}

func (button *Button) GetSize() (int, int) {
	return button.w, button.h
}

func (button *Button) draw() error {
	px, py := button.parent.GetAbsolutePosition()

	if button.is_pressed {
		screen.DrawScaledImageAt(button.pressed_img, button.x+px, button.y+py, button.w, button.h)
	} else {
		screen.DrawScaledImageAt(button.img, button.x+px, button.y+py, button.w, button.h)
	}
	return nil
}
