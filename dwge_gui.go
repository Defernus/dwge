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

type GuiContainer interface {
	AddElement(element GuiElement) (relative_id int, err error)
	GetElement(relative_id int) (element GuiElement, err error)
	PopElement(relative_id int) (element GuiElement, err error)
}
