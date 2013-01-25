package Game

import (
	"github.com/vova616/GarageEngine/Engine"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

type Chud struct {
	Engine.BaseComponent
}

func NewChud() *Chud {
	return &Chud{Engine.NewComponent()}
}
