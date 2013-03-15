package Player

import (
	// "github.com/mentaman/PirateLand/Game/GUI"
	"github.com/vova616/garageEngine/engine"
	// "github.com/vova616/garageEngine/engine/components"
	"math/rand"
	// "github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

const (
	s_none   int = -1
	s_rand   int = 0
	s_hp     int = 2
	s_strong int = 3
)

var (
	SkillsAtlas  *engine.ManagedAtlas
	CurrentSkill = 0
	Sk           *SkillManager
)

const (
	Spr_skillGUI    = 1
	Spr_skillHp     = 2
	Spr_skillStrong = 3
)

type SkillManager struct {
	engine.BaseComponent
	Skills []*Skill
}

func (s *SkillManager) AddSkill(ty int) {
	if ty == s_rand {
		ty = rand.Int()%2 + 2
	}
	SkillObj = engine.NewGameObject("Skill")
	SkillObj.AddComponent(engine.NewSprite2(SkillsAtlas.Texture, engine.IndexUV(SkillsAtlas, ty)))
	sk := SkillObj.AddComponent(NewSkill(ty)).(*Skill)
	s.Skills = append(s.Skills, sk)
	sk.SetPlace(len(s.Skills))
	SkillObj.AddComponent(engine.NewPhysics(false))
	SkillObj.Physics.Shape.IsSensor = true
	SkillObj.Physics.Body.IgnoreGravity = true
	sk.Transform().SetDepth(2)
	sk.Transform().SetScalef(30, 30)
	sk.Transform().SetParent(s.Transform().Parent())
}
func NewSkillManager() *SkillManager {
	return &SkillManager{engine.NewComponent(), nil}
}
