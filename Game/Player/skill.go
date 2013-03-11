package Player

import (
	// "github.com/mentaman/PirateLand/Game/GUI"
	"github.com/vova616/garageEngine/engine"
	// "github.com/vova616/garageEngine/engine/components"

	// "github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

var (
	SkillObj *engine.GameObject
)

type Skill struct {
	engine.BaseComponent
	value  float32
	place  int
	skType int
}

func (s *Skill) SetPlace(place int) {
	p := s.Transform().Position()
	s.place = place
	p.Y = -300
	p.X = float32(place)*50 - 500
	s.Transform().SetPosition(p)
	if s.Transform().Position().X > 500 {
		s.GameObject().Sprite.Color.A = 0
	}

}
func NewSkill(ty int) *Skill {
	return &Skill{engine.NewComponent(), 0, 0, ty}
}
