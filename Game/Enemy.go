package Game

import (
	"github.com/vova616/garageEngine/engine"
	"math"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

type Enemy struct {
	engine.BaseComponent

	frames    int
	Hp        float32
	MaxHp     float32
	HpB       *Bar
	OnGround  bool
	Attack    bool
	able      bool
	jump      bool
	hit       bool
	hitable   bool
	speed     float32
	width     float32
	height    float32
	jumppower float32
	LastFloor *engine.GameObject

	Hitted *engine.Coroutine
}

func NewEnemy(Hp *Bar) *Enemy {
	return &Enemy{engine.NewComponent(), 0, 100, 100, Hp, false, false, true, false, false, true, 60, 0, 0, 3000, nil, engine.StartCoroutine(func() {})}

}
func (s *Enemy) Start() {
	s.GameObject().Sprite.SetAnimation("enemy_jump")

	s.GameObject().Physics.Body.SetMoment(engine.Inf)
	s.width = s.Transform().WorldScale().X
	s.height = s.Transform().WorldScale().Y
}
func (s *Enemy) Update() {
	ph := s.GameObject().Physics.Body
	s.GameObject().Sprite.SetAlign(engine.AlignTopCenter)
	if float32(math.Abs(float64(ph.Velocity().X))) > 3 {
		if s.GameObject().Sprite.CurrentAnimation() == "enemy_stand" {
			s.GameObject().Sprite.SetAnimation("enemy_walk")
		}
	} else if !s.Attack {
		s.GameObject().Sprite.SetAnimation("enemy_stand")
	}
	if s.OnGround == false {
		s.frames++
	} else {
		s.frames = 0
	}
	d := s.Transform().WorldPosition()
	s.HpB.Transform().SetWorldPosition(d.Add(engine.NewVector2(0, 30)))
	if s.able {
		if plComp.Transform().WorldPosition().X > s.Transform().WorldPosition().X {
			ph.AddForce(s.speed, 0)
			s.GameObject().Transform().SetScalef(s.width, s.height)

		} else {
			ph.AddForce(-s.speed, 0)
			s.GameObject().Transform().SetScalef(-s.width, s.height)
		}
		d = plComp.Transform().WorldPosition()
		if d.Distance(s.Transform().WorldPosition()) < 50 {
			s.Attack = true

			s.GameObject().Sprite.SetAnimation("enemy_attack")
			s.GameObject().Sprite.AnimationEndCallback = func(sprite *engine.Sprite) {
				s.Attack = false
				s.GameObject().Sprite.SetAnimation("enemy_stand")
			}

		}
		if s.jump && s.OnGround {
			s.GameObject().Physics.Body.AddForce(0, s.jumppower)
		}
	}
	if s.hit {
		s.GameObject().Sprite.SetAnimation("enemy_hit")
	}
	if (s.GameObject().Sprite.CurrentAnimation() == "enemy_stand" || s.GameObject().Sprite.CurrentAnimation() == "enemy_walk") && !s.OnGround && s.frames > 15 {
		s.GameObject().Sprite.SetAnimation("enemy_jump")
	}

}
func (s *Enemy) OnCollisionPostSolve(arbiter engine.Arbiter) {
	if arbiter.GameObjectB().Tag != "lader" && arbiter.GameObjectA().Tag != "lader" {
		count := 0
		for _, con := range arbiter.Contacts {
			if arbiter.Normal(con).Y < -0.9 {
				count++

			}
		}
		if count >= 1 {
			if s.GameObject().Sprite.CurrentAnimation() == "enemy_jump" {
				s.GameObject().Sprite.SetAnimation("enemy_stand")
			}
			s.LastFloor = arbiter.GameObjectB()

			s.OnGround = true
		}
		count = 0
		for _, con := range arbiter.Contacts {
			if math.Abs(float64(arbiter.Normal(con).X)) > 0.9 {
				count++

			}
		}
		if count >= 1 {
			s.jump = true
		}
		if arbiter.GameObjectB().Tag == "player" {
			if plComp.hitable && s.Attack {
				plComp.Hit()
			}
			if s.hitable && plComp.Attack {
				s.Hit()
			}
		}
	}

}
func (s *Enemy) FixedUpdate() {
	s.OnGround = false
	s.LastFloor = nil
	s.jump = false
}
func (s *Enemy) OnDestroy() {
	s.HpB.GameObject().Destroy()
}
func (s *Enemy) Hit() {
	if s.Hitted.State == engine.Ended {
		s.Hitted = engine.StartCoroutine(func() {
			s.hit = true
			s.Attack = false
			s.LastFloor = nil
			s.hitable = false
			s.able = false
			if s.GameObject() != nil {
				s.GameObject().Physics.Body.AddForce(0, s.jumppower)
			}
			s.SubLife(50)
			engine.CoSleep(3)
			s.hit = false
			s.able = true
			s.hitable = true
			if s.GameObject() != nil {
				s.GameObject().Sprite.SetAnimation("enemy_stand")
			}
		})

	}
}
func (s *Enemy) SubLife(hp float32) {
	s.Hp -= hp
	if s.Hp < 0 {
		s.GameObject().Destroy()
	}
	s.HpB.SetValue(s.Hp, s.MaxHp)
}
