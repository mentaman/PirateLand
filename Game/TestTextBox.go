package Game

import (
	"github.com/vova616/GarageEngine/Engine/Components"
	"github.com/vova616/GarageEngine/Engine/Input"
)

type TestTextBox struct {
	*Components.UIText
	able bool
}

func NewTestBox() *TestTextBox {
	return &TestTextBox{Components.NewUIText(ArialFont2, "var: "), false}
}
func (s *TestTextBox) Update() {
	if Input.KeyPress(Input.KeyF2) {
		s.able = !s.able
		s.SetFocus(s.able)
		s.SetWritable(s.able)
	}
}
