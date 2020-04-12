package dwge

import (
	"fmt"
	"sync"
)

type MainWindow struct {
	mx             sync.Mutex
	selected_scene int
	scenes         [][]GuiElement
}

func newMainWindow() *MainWindow {
	return &MainWindow{
		scenes: [][]GuiElement{make([]GuiElement, 0)},
	}
}

func (mwin *MainWindow) GetCurentSceneId() int {
	return mwin.selected_scene
}

func (mwin *MainWindow) CreateNewScene() int {
	mwin.scenes = append(mwin.scenes, make([]GuiElement, 0))
	return len(mwin.scenes) - 1
}

//TODO fix bug when scene changed while passing throw the element of scene
func (mwin *MainWindow) SetScene(id int) {
	for i := 0; i != 5; i++ {
		mwin.mouseButtonUpEvent(mouse_x, mouse_y, i)
	}
	mwin.selected_scene = id
}

func (mwin *MainWindow) draw() error {
	for i := range mwin.scenes[mwin.selected_scene] {
		if err := mwin.scenes[mwin.selected_scene][i].draw(); err != nil {
			return err
		}
	}
	return nil
}

func (mwin *MainWindow) mouseButtonDownEvent(x, y, button int) bool {
	fmt.Println(button)
	for _, v := range mwin.scenes[mwin.selected_scene] {
		switch v.(type) {
		case GuiClickable:
			if v.(GuiClickable).mouseButtonDownEvent(x, y, button) {
				return true
			}
		}
	}
	return true
}

func (mwin *MainWindow) mouseButtonUpEvent(x, y, button int) {
	for _, v := range mwin.scenes[mwin.selected_scene] {
		switch v.(type) {
		case GuiClickable:
			v.(GuiClickable).mouseButtonUpEvent(x, y, button)
		}
	}
}

func (mwin *MainWindow) GetParent() GuiElement {
	return nil
}

func (mwin *MainWindow) setParent(parent GuiElement) error {
	return fmt.Errorf("can not set parent to mainWindow")
}

func (mwin *MainWindow) clearParent() {}

func (mwin *MainWindow) GetRelativePosition() (int, int) {
	return 0, 0
}

func (mwin *MainWindow) GetAbsolutePosition() (int, int) {
	return 0, 0
}

func (mwin *MainWindow) GetSize() (int, int) {
	return screen.w, screen.h
}

func (mwin *MainWindow) AddElement(element GuiElement) error {
	if err := element.setParent(mwin); err != nil {
		return err
	}
	mwin.scenes[mwin.selected_scene] = append(mwin.scenes[mwin.selected_scene], element)
	return nil
}

func (mwin *MainWindow) RemoveElement(element GuiElement) error {
	for i, v := range mwin.scenes[mwin.selected_scene] {
		if v == element {
			if i == 0 {
				mwin.scenes[mwin.selected_scene] = mwin.scenes[mwin.selected_scene][1:]
			} else if i == len(mwin.scenes[mwin.selected_scene])-1 {
				mwin.scenes[mwin.selected_scene] = mwin.scenes[mwin.selected_scene][:len(mwin.scenes[mwin.selected_scene])-1]
			} else if i > len(mwin.scenes[mwin.selected_scene])/2 {
				copy(mwin.scenes[mwin.selected_scene][i:], mwin.scenes[mwin.selected_scene][i+1:])
				mwin.scenes[mwin.selected_scene] = mwin.scenes[mwin.selected_scene][:len(mwin.scenes[mwin.selected_scene])-1]
			} else {
				copy(mwin.scenes[mwin.selected_scene][1:i+1], mwin.scenes[mwin.selected_scene][0:i])
				mwin.scenes[mwin.selected_scene] = mwin.scenes[mwin.selected_scene][1:]
			}
			return nil
		}
	}
	return fmt.Errorf("could not find element")
}
