package Player

import (
	// "github.com/mentaman/PirateLand/Game/GUI"
	"github.com/vova616/garageEngine/engine"
	// "github.com/vova616/garageEngine/engine/components"

	// "github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

const (
	s_none   int = -1
	s_strong int = 0
)

var (
	SkillsAtlas  *engine.ManagedAtlas
	CurrentSkill = 0
	Sk           *SkillManager
)

const (
	Spr_skillGUI = 1
)

type SkillManager struct {
	engine.BaseComponent
	Skills []Skill
}

func (s *SkillManager) AddSkill(kind int) {

}
func NewSkillManager() *SkillManager {
	return &SkillManager{engine.NewComponent(), nil}
}
