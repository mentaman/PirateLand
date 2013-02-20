package GUI

import (
	"github.com/vova616/garageEngine/engine"
	"github.com/vova616/garageEngine/engine/components"

	"github.com/mentaman/PirateLand/Game/Fonts"
	"strconv"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

type Bar struct {
	engine.BaseComponent
	inited bool
	value  float32
	width  float32
	label  *engine.GameObject
	text   *components.UIText
}

func NewBar(width float32) *Bar {
	return &Bar{engine.NewComponent(), false, 0, width, nil, nil}
}
func (s *Bar) Start() {
	if !s.inited {
		label := engine.NewGameObject("Label")
		label.Transform().SetParent(s.Transform().Parent())
		p := s.Transform().WorldPosition()
		si := s.Transform().WorldScale()
		label.Transform().SetWorldPositionf(p.X+si.X/2, p.Y+si.Y/2)
		label.Transform().SetScalef(20, 20)
		txt2 := label.AddComponent(components.NewUIText(Fonts.ArialFont2, "100/100")).(*components.UIText)
		txt2.SetAlign(engine.AlignCenter)
		txt2.Transform().SetDepth(s.Transform().Depth() + 1)
		s.text = txt2
		s.label = label
	}

	s.inited = true

}
func (s *Bar) GetValue() float32 {
	return s.value
}
func (s *Bar) Update() {
	p := s.Transform().WorldPosition()
	si := s.Transform().WorldScale()
	s.label.Transform().SetWorldPositionf(p.X+si.X/2, p.Y+si.Y/2)
}
func (s *Bar) OnDestroy() {
	if s.label.GameObject() != nil {
		s.label.GameObject().Destroy()
	}
}
func (s *Bar) SetValue(min, max float32) {
	s.value = min / max
	sc := s.GameObject().Transform().Scale()
	sc.X = s.width * s.value
	s.text.SetString(strconv.Itoa(int(min)) + "/" + strconv.Itoa(int(max)))
	s.GameObject().Transform().SetScale(sc)
}
