package Game

import (
	// "github.com/vova616/GarageEngine/Engine"
	"github.com/vova616/GarageEngine/Engine/Components"
	"github.com/vova616/GarageEngine/Engine/Input"
	"strconv"
)

type TestTextBox struct {
	*Components.UIText
	able bool
	Do   func(*TestTextBox)
	V    float64
}

func NewTestBox(do func(*TestTextBox)) *TestTextBox {
	return &TestTextBox{Components.NewUIText(ArialFont2, ""), false, do, 0}
}
func (s *TestTextBox) Update() {
	if Input.KeyPress(Input.KeyF2) {
		s.able = !s.able
		s.SetFocus(s.able)
		s.SetWritable(s.able)
	}
	if s.able && Input.KeyPress(Input.KeyEnter) {
		s.V, _ = strconv.ParseFloat(s.String(), 64)
		s.Do(s)
	}

	s.UIText.Update()
}
