package Player

import (
	// "github.com/mentaman/PirateLand/Game/GUI"
	"github.com/vova616/garageEngine/engine"
	// "github.com/vova616/garageEngine/engine/components"

	// "github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

type Skill struct {
	engine.BaseComponent
	x, y, value float32
	skType      int
}

func CreateSkill() {

}
func (s *Skill) SetPlace(place int) {

}
func NewSkill() *Skill {
	return &Skill{engine.NewComponent(), 0, 0, 0, s_strong}
}
