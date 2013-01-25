package Game

import (
	"github.com/vova616/GarageEngine/Engine"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

type Chud struct {
	Engine.BaseComponent
	hp  *Bar
	cp  *Bar
	exp *Bar
}

func NewChud() *Chud {
	return &Chud{Engine.NewComponent(), nil, nil, nil}
}
func (s *Chud) Start() {
	s.GameObject().Sprite.SetAlign(Engine.AlignBottomLeft)
	hpB := Engine.NewGameObject("hpBar")
	hpB.Transform().SetParent2(up)
	hpB.AddComponent(NewBar(1.0, atlas, spr_chudHp, 20, 20, Engine.Vector{380, 550, 0}))

}
