package Game

import (
	// "github.com/vova616/garageEngine/Engine"
	"github.com/vova616/garageEngine/engine/components"
	"github.com/vova616/garageEngine/engine/input"
	"strconv"
)

type TestTextBox struct {
	*components.UIText
	able bool
	Do   func(*TestTextBox)
	V    float64
}

func NewTestBox(do func(*TestTextBox)) *TestTextBox {
	return &TestTextBox{components.NewUIText(ArialFont2, ""), false, do, 0}
}
func (s *TestTextBox) Update() {
	if input.KeyPress(input.KeyF2) {
		s.able = !s.able
		s.SetFocus(s.able)
		s.SetWritable(s.able)
	}
	if s.able && input.KeyPress(input.KeyEnter) {
		s.V, _ = strconv.ParseFloat(s.String(), 64)
		s.Do(s)
	}

	s.UIText.Update()
}
