package dwge

import (
	"fmt"
)

type mainWindow struct {
	last_id  int
	elements map[int]GuiElement
}

func newMainWindow() *mainWindow {
	return &mainWindow{
		last_id:  0,
		elements: make(map[int]GuiElement),
	}
}

func (mwin *mainWindow) draw() error {
	for i := range mwin.elements {
		if err := mwin.elements[i].draw(); err != nil {
			return err
		}
	}
	return nil
}

func (mwin *mainWindow) GetParent() GuiElement {
	return nil
}

func (mwin *mainWindow) setParent(parent GuiElement) error {
	return fmt.Errorf("can not set parent to mainWindow")
}

func (mwin *mainWindow) clearParent() {}

func (mwin *mainWindow) GetRelativePosition() (int, int) {
	return 0, 0
}

func (mwin *mainWindow) GetAbsolutePosition() (int, int) {
	return 0, 0
}

func (mwin *mainWindow) GetSize() (int, int) {
	return screen.w, screen.h
}

func (mwin *mainWindow) AddElement(element GuiElement) (int, error) {
	if err := element.setParent(mwin); err != nil {
		return -1, err
	}
	mwin.elements[mwin.last_id] = element
	mwin.last_id++
	return mwin.last_id - 1, nil
}

func (mwin *mainWindow) GetElement(relative_id int) (GuiElement, error) {
	if el, exist := mwin.elements[relative_id]; exist {
		return el, nil
	}
	return nil, fmt.Errorf("could not find element with id %v", relative_id)
}

func (mwin *mainWindow) PopElement(relative_id int) (GuiElement, error) {
	if el, exist := mwin.elements[relative_id]; exist {
		el.clearParent()
		delete(mwin.elements, relative_id)
		return el, nil
	}
	return nil, fmt.Errorf("could not find element with id %v", relative_id)
}
