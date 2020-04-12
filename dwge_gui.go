package dwge

type GuiElement interface {
	GetParent() (parent GuiElement)
	setParent(parent GuiElement) error
	clearParent()

	GetRelativePosition() (x, y int)
	GetAbsolutePosition() (x, y int)
	GetSize() (width, height int)

	draw() error
}

type GuiClickable interface {
	mouseButtonDownEvent(x, y, button int) bool
	mouseButtonUpEvent(x, y, button int)
}

type GuiContainer interface {
	AddElement(element GuiElement) error
	RemoveElement(element GuiElement) error

	mouseButtonDownEvent(x, y, button int) bool
	mouseButtonUpEvent(x, y, button int)
}
