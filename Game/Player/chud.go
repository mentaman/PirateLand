package Player

import (
	"github.com/mentaman/PirateLand/Game/GUI"
	"github.com/vova616/garageEngine/engine"
	"github.com/vova616/garageEngine/engine/components"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

var (
	Ch        *Chud
	ChudAtlas *engine.ManagedAtlas
)

const (
	Spr_chud    = 1
	Spr_chudHp  = 2
	Spr_chudCp  = 3
	Spr_chudExp = 4
	Spr_scroll  = 5
)

type Chud struct {
	engine.BaseComponent
	Scrolls float32
	Hp      *GUI.Bar
	Cp      *GUI.Bar
	Exp     *GUI.Bar
	Money   *components.UIText
	Level   *components.UIText
}

func CreateChud() {
	chud := engine.NewGameObject("chud")
	Ch = chud.AddComponent(NewChud()).(*Chud)
	chud.AddComponent(engine.NewSprite2(ChudAtlas.Texture, engine.IndexUV(ChudAtlas, Spr_chud)))
	chud.Transform().SetWorldPositionf(-440, 190)
	chud.Transform().SetWorldScalef(100, 100)
	chud.Transform().SetDepth(1)
}
func (s *Chud) AddScroll() {
	s.Scrolls++
	scr := Scroll.Clone()
	scr.Transform().SetWorldPosition(engine.Vector{s.Scrolls*20 + -440, 190, 0})
	scr.Transform().SetParent(s.Transform().Parent())
}
func NewChud() *Chud {
	return &Chud{engine.NewComponent(), 0, nil, nil, nil, nil, nil}
}
