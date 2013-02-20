package Enemy

import (
	"github.com/mentaman/PirateLand/Game/GUI"
	"github.com/vova616/garageEngine/engine"

	"github.com/mentaman/PirateLand/Game/Objects"
	"github.com/mentaman/PirateLand/Game/Player"
	//"log"
	"math"
	"math/rand"
	"strconv"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

type Enemy struct {
	engine.BaseComponent

	frames    int
	Hp        float32
	MaxHp     float32
	HpB       *GUI.Bar
	OnGround  bool
	Attack    bool
	able      bool
	jump      bool
	hit       bool
	hitable   bool
	isClose   bool
	strengh   int
	speed     float32
	width     float32
	height    float32
	jumppower float32
	LastFloor *engine.GameObject
	target    engine.Vector
	Hitted    *engine.Coroutine
}

var (
	Regular *engine.GameObject
	Atlas   *engine.ManagedAtlas
	HpBar   *engine.GameObject
)

func CreateEnemy() {
	uvs, ind := engine.AnimatedGroupUVs(Atlas, "enemy_walk", "enemy_stand", "enemy_attack", "enemy_jump", "enemy_hit")
	Regular = engine.NewGameObject("Enemy")
	Regular.AddComponent(engine.NewSprite3(Atlas.Texture, uvs))
	Regular.Sprite.BindAnimations(ind)
	Regular.Transform().SetWorldScalef(50, 100)
	Regular.AddComponent(engine.NewPhysics(false, 1, 1))
	Regular.Physics.Shape.SetFriction(0.7)
	Regular.Physics.Shape.SetElasticity(0.2)

	HpBar = engine.NewGameObject("hpBar")
	HpBar.GameObject().AddComponent(engine.NewSprite2(Player.ChudAtlas.Texture, engine.IndexUV(Player.ChudAtlas, Player.Spr_chudHp)))
	HpBar.GameObject().Sprite.SetAlign(engine.AlignLeft)
	HpBar.GameObject().Transform().SetWorldScalef(10, 15)
	HpBar.AddComponent(GUI.NewBar(10))
}
func NewEnemy(Hp *GUI.Bar) *Enemy {
	return &Enemy{engine.NewComponent(), 0, 100, 100, Hp, false, false, true, false, false, true, false, 5, 60, 0, 0, 3000, nil, engine.Vector{0, 0, 0}, engine.StartCoroutine(func() {})}

}
func (s *Enemy) Start() {
	s.GameObject().Sprite.SetAnimation("enemy_jump")

	s.GameObject().Physics.Body.SetMoment(engine.Inf)
	s.width = s.Transform().WorldScale().X
	s.height = s.Transform().WorldScale().Y
}
func (s *Enemy) Update() {
	plpc := Player.Pl.Transform().WorldPosition()
	if plpc.Distance(s.Transform().WorldPosition()) < 200 {
		s.target = plpc
		s.speed = 60
		s.isClose = false
	} else {
		if !s.isClose && s.speed != 0 {
			s.target = s.Transform().WorldPosition()
			r := float32(rand.Int()%200) - 100
			if r > 0 {
				r += 20
			} else {
				r -= 20
			}
			s.target = s.target.Add(engine.NewVector2(r, 0))

			s.isClose = true
		}
		if s.isClose {
			if s.target.Distance(s.Transform().WorldPosition()) < 60 {

				s.isClose = false
				tr := s.Transform()
				engine.StartCoroutine(func() {
					s.speed = 0
					engine.CoSleep(float32(rand.Int()%2) + 3)
					if plpc.Distance(tr.WorldPosition()) >= 200 {
						s.speed = 60
						s.isClose = true
						r := float32(rand.Int()%200) - 100
						if r > 0 {
							r += 60
						} else {
							r -= 60
						}
						s.target = tr.WorldPosition()
						s.target = s.target.Add(engine.NewVector2(r, 0))
					}
				})

				return
			}
		}
	}
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
	s.HpB.Transform().SetWorldPosition(d.Add(engine.NewVector2(-50, -10)))
	if s.able {
		if s.target.X > s.Transform().WorldPosition().X {
			ph.AddForce(s.speed, 0)
			s.GameObject().Transform().SetScalef(s.width, s.height)

		} else {
			ph.AddForce(-s.speed, 0)
			s.GameObject().Transform().SetScalef(-s.width, s.height)
		}
		d = Player.PlComp.Transform().WorldPosition()
		if d.Distance(s.Transform().WorldPosition()) < 50 && !s.Attack {
			s.Attack = true

			s.GameObject().Sprite.SetAnimation("enemy_attack")
			s.GameObject().Sprite.AnimationEndCallback = func(sprite *engine.Sprite) {
				s.Attack = false
				s.GameObject().Sprite.SetAnimation("enemy_stand")
				s.GameObject().Sprite.AnimationEndCallback = nil
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
	if arbiter.GameObjectB() != nil {
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
				if Player.PlComp.Hitable && s.Attack {
					Player.PlComp.Hit(rand.Int()%s.strengh + s.strengh)
				}
				if s.hitable && Player.PlComp.Attack {
					s.Hit(int(float32(Player.PlComp.Strengh*(rand.Int()%4+7)) / 10))
				}
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
	Player.PlComp.AddExp(float32(rand.Int()%20) + 20)
	if rand.Int()%100 > 70 {
		it := Objects.RandomItem().Clone()
		it.Transform().SetPosition(s.Transform().Position())
		it.Transform().SetParent(s.Transform().Parent())
		it.ComponentTypeOfi((*Objects.Item).TypeOf(nil)).(*Objects.Item).Pop()
	}
}
func (s *Enemy) Hit(dmg int) {
	if s.Hitted.State == engine.Ended {
		s.Hitted = engine.StartCoroutine(func() {
			GUI.NewUpTextObj(strconv.Itoa(dmg), s.Transform(), 20)
			s.hit = true
			s.Attack = false
			s.LastFloor = nil
			s.hitable = false
			s.able = false
			if s.GameObject() != nil {
				s.GameObject().Physics.Body.AddForce(0, s.jumppower)
			}
			s.SubLife(float32(dmg))
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
