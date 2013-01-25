package Game

import (
	"github.com/vova616/GarageEngine/Engine"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

type Bar struct {
	Engine.BaseComponent
	value    float32
	myAtlas  *Engine.ManagedAtlas
	sprInd   int
	width    float32
	height   float32
	position Engine.Vector
}

func NewBar(start_value float32, atl *Engine.ManagedAtlas, sprI int, width, height float32, pos Engine.Vector) *Bar {
	return &Bar{Engine.NewComponent(), start_value, atl, sprI, width, height, pos}
}
func (s *Bar) Start() {
	s.GameObject().AddComponent(Engine.NewSprite2(s.myAtlas.Texture, Engine.IndexUV(s.myAtlas, s.sprInd)))
	s.GameObject().Transform().SetWorldPosition(s.position)
	s.GameObject().Sprite.SetAlign(Engine.AlignBottomLeft)
	s.GameObject().Transform().SetWorldScalef(s.width*s.value, s.height)

}
func (s *Bar) GetValue() float32 {
	return s.value
}
func (s *Bar) SetValue(val float32) {
	s.value = val
	s.GameObject().Transform().SetWorldScalef(s.width*s.value, s.height)
}
