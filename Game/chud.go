package Game

import (
	"github.com/vova616/garageEngine/engine"
	"github.com/vova616/garageEngine/engine/components"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

type Chud struct {
	engine.BaseComponent
	Hp    *Bar
	Cp    *Bar
	Exp   *Bar
	Money *components.UIText
	Level *components.UIText
}

func NewChud() *Chud {
	return &Chud{engine.NewComponent(), nil, nil, nil, nil, nil}
}
