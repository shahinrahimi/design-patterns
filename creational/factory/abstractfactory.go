package factory

import "fmt"

// AbstractProduct: Button
type Button interface {
  Render() string  
}

// AbstractProduct: TextBox
type TextBox interface {
  Render() string
} 

// ConcreteProducts for MacOS
type MacButton struct{}
type MacTextBox struct{}

func (b *MacButton) Render() string {
  return "Rendering a MacOS Button"
}

func (b *MacTextBox) Render() string {
  return "Rendering a MacOS TextBox"
}

// ConcreteProducts for Windows
type WinButton struct{}
type WinTextBox struct{}

func (b *WinButton) Render() string {
  return "Rendering a MacOS Button"
}

func (b *WinTextBox) Render() string {
  return "Rendering a MacOS TextBox"
}


// AbstractFactory
type GUIFactory interface {
  CreateButton() Button
  CreateTextBox() TextBox
}

// ConcreteFactory for MacOS
type MacFactory struct{}

func (f *MacFactory) CreateButton() Button {
  return &MacButton{}
}

func (f *MacFactory) CreateTextBox() TextBox {
  return &MacTextBox{}
}

// ConcreteFactory for Windows
type WinFactory struct{}

func (f *WinFactory) CreateButton() Button {
  return &WinButton{}
}

func (f *WinFactory) CreateTextBox() TextBox {
  return &WinTextBox{}
}


func RunAbstractFactory() {
  var f GUIFactory
  
  // choose which factory to use on OS or configuration
  f = &MacFactory{}

  // create UI components
  button := f.CreateButton()
  textBox := f.CreateTextBox()

  fmt.Println(button.Render())
  fmt.Println(textBox.Render())
}
