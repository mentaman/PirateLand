package Player

import (
	// "github.com/mentaman/PirateLand/Game/GUI"
	"github.com/vova616/garageEngine/engine"
	// "github.com/vova616/garageEngine/engine/components"

	"github.com/vova616/garageEngine/engine/input"
	// "github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

var (
	SkillObj *engine.GameObject
	selected int = 0
	count    int = 0
)

type Skill struct {
	engine.BaseComponent
	value  float32
	place  int
	skType int
	mouse  bool
	ID     int
}

func (s *Skill) GetPlace() int {
	return s.place
}

func (s *Skill) SetPlace(place int) {
	p := s.Transform().Position()
	s.place = place
	p.Y = -300

	p.X = float32(place)*50 - 440
	s.Transform().SetPosition(p)
	if s.Transform().Position().X > 400 {
		s.GameObject().Sprite.Color.A = 0
	}

}
func (s *Skill) Update() {
	if selected == s.ID {
		s.GameObject().Sprite.Color = engine.Color{1, 1, 1, 1}
	} else if s.mouse {
		s.GameObject().Sprite.Color = engine.Color{0.4, 0.4, 0.7, 1}
	} else {
		s.GameObject().Sprite.Color = engine.Color{0.2, 0.2, 0.2, 1}
	}
	if s.mouse && input.MousePress(input.MouseLeft) {
		selected = s.ID
	}

	if s.mouse && input.MousePress(input.MouseRight) {
		for _, sk := range Sk.Skills {
			p := sk.GetPlace()
			if p > s.place {
				sk.SetPlace(p - 1)
			}
		}
		s.SetPlace(len(Sk.Skills))
	}
}
func (s *Skill) OnMouseEnter(col engine.Arbiter) bool {
	s.mouse = true
	return false
}
func (s *Skill) OnMouseExit(col engine.Arbiter) {
	s.mouse = false
}

func NewSkill(ty int) *Skill {
	count++
	return &Skill{engine.NewComponent(), 0, 0, ty, false, count}
}
