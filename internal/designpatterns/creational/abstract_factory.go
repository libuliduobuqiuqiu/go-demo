package creational

import "fmt"

type Button interface {
	Click()
}

type CheckBox interface {
	Check()
}

type UiFactory interface {
	CreateButton() Button
	CreateCheckBox() CheckBox
}

type WinButton struct{}

func (w WinButton) Click() {
	fmt.Println("Win Button")
}

type WinCheckBox struct{}

func (w WinCheckBox) Check() {
	fmt.Println("Win CheckBox")
}

type WinFactory struct{}

func (w WinFactory) CreateButton() Button {
	return WinButton{}
}

func (w WinFactory) CreateCheckBox() CheckBox {
	return WinCheckBox{}
}

type MacButton struct{}

func (m MacButton) Click() {
	fmt.Println("Mac Button")
}

type MacCheckBox struct{}

func (m MacCheckBox) Check() {
	fmt.Println("Mac CheckBox")
}

type MacFactory struct{}

func (m MacFactory) CreateButton() Button {
	return MacButton{}
}

func (m MacFactory) CreateCheckBox() CheckBox {
	return MacCheckBox{}
}

// Abstract Factory:  The abstract factory patterns is used to creaete multiple related products.
// Compare to the factory patterns, add new products only requires implementing a new factory.
func Render(f UiFactory) {
	b := f.CreateButton()
	c := f.CreateCheckBox()

	b.Click()
	c.Check()
}
