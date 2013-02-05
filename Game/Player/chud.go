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
)

type Chud struct {
	engine.BaseComponent
	Hp    *GUI.Bar
	Cp    *GUI.Bar
	Exp   *GUI.Bar
	Money *components.UIText
	Level *components.UIText
}

func CreateChud() {
	chud := engine.NewGameObject("chud")
	Ch = chud.AddComponent(NewChud()).(*Chud)
	chud.AddComponent(engine.NewSprite2(ChudAtlas.Texture, engine.IndexUV(ChudAtlas, Spr_chud)))
	chud.Transform().SetWorldPositionf(200, 550)
	chud.Transform().SetWorldScalef(100, 100)
}

func NewChud() *Chud {
	return &Chud{engine.NewComponent(), nil, nil, nil, nil, nil}
}
