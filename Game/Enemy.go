package Game

import (
	"github.com/vova616/GarageEngine/Engine"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

type Enemy struct {
	Engine.BaseComponent
	Hp    float32
	MaxHp float32
	HpB   *Bar
}

func NewEnemy(Hp *Bar) *Enemy {
	return &Enemy{Engine.NewComponent(), 0, 100, Hp}
}
func (s *Enemy) Update() {
	d := s.Transform().WorldPosition()
	s.HpB.Transform().SetWorldPosition(d.Add(Engine.NewVector2(0, 20)))
}
