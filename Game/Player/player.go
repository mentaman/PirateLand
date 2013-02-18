package Player

import (
	"github.com/mentaman/PirateLand/Game/GUI"

	"github.com/mentaman/PirateLand/Game/Fonts"
	"github.com/vova616/garageEngine/engine"
	"github.com/vova616/garageEngine/engine/components"
	"github.com/vova616/garageEngine/engine/input"
	"math"
	"strconv"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

var (
	PlComp *Player
	Pl     *engine.GameObject
	Atlas  *engine.ManagedAtlas
	Scroll *engine.GameObject
)

type Player struct {
	engine.BaseComponent
	frames    int
	money     int
	level     int
	Strengh   int
	Exp       float32
	MaxExp    float32
	Cp        float32
	MaxCp     float32
	Hp        float32
	MaxHp     float32
	width     float32
	height    float32
	speed     float32
	jumpPower float32
	Attack    bool

	OnGround  bool
	bend      bool
	right     float32
	able      bool
	hit       bool
	Hitable   bool
	pLader    *engine.GameObject
	pSplint   *engine.GameObject
	LastFloor *engine.GameObject
	Hitted    *engine.Coroutine

	MenuScene func()
}

const (
	stand_height = 100
	bend_height  = 70
)

func CreatePlayer() {
	uvs, ind := engine.AnimatedGroupUVs(Atlas, "player_walk", "player_stand", "player_attack", "player_jump", "player_bend", "player_hit", "player_climb")
	Pl = engine.NewGameObject("Player")
	Pl.AddComponent(engine.NewSprite3(Atlas.Texture, uvs))
	Pl.Sprite.BindAnimations(ind)
	Pl.Sprite.AnimationSpeed = 10
	Pl.Transform().SetWorldPositionf(100, 180)
	Pl.Transform().SetWorldScalef(50, 100)
	PlComp = Pl.AddComponent(NewPlayer()).(*Player)
	Pl.AddComponent(components.NewSmoothFollow(nil, 1, 30))
	Pl.AddComponent(engine.NewPhysics(false, 1, 1))
	Pl.Physics.Shape.SetFriction(0.7)
	Pl.Physics.Shape.SetElasticity(0.2)
	Pl.Tag = "player"

	Hp := engine.NewGameObject("hpBar")
	Hp.Transform().SetDepth(2)
	Hp.GameObject().AddComponent(engine.NewSprite2(ChudAtlas.Texture, engine.IndexUV(ChudAtlas, Spr_chudHp)))
	Hp.GameObject().Sprite.SetAlign(engine.AlignLeft)
	Hp.GameObject().Transform().SetWorldPosition(engine.Vector{156, 580, 0})
	Hp.GameObject().Transform().SetWorldScalef(16.2, 20)
	Ch.Hp = (Hp.AddComponent(GUI.NewBar(17))).(*GUI.Bar)

	Cp := engine.NewGameObject("cpBar")
	Cp.Transform().SetDepth(2)
	Cp.GameObject().AddComponent(engine.NewSprite2(ChudAtlas.Texture, engine.IndexUV(ChudAtlas, Spr_chudCp)))
	Cp.GameObject().Sprite.SetAlign(engine.AlignLeft)
	Cp.GameObject().Transform().SetWorldPosition(engine.Vector{156, 555, 0})
	Cp.GameObject().Transform().SetWorldScalef(15.05, 20)
	Ch.Cp = (Cp.AddComponent(GUI.NewBar(17))).(*GUI.Bar)

	Scroll = engine.NewGameObject("scroll")
	Scroll.AddComponent(engine.NewSprite2(Atlas.Texture, engine.IndexUV(Atlas, Spr_scroll)))
	Scroll.Transform().SetWorldScalef(20, 20)
	Scroll.Transform().SetDepth(2)

	Exp := engine.NewGameObject("expBar")
	Exp.GameObject().AddComponent(engine.NewSprite2(ChudAtlas.Texture, engine.IndexUV(ChudAtlas, Spr_chudExp)))
	Exp.GameObject().Sprite.SetAlign(engine.AlignLeft)
	Exp.Transform().SetDepth(2)
	Exp.GameObject().Transform().SetWorldPosition(engine.Vector{156, 530, 0})
	Exp.GameObject().Transform().SetWorldScalef(17, 20)
	Ch.Exp = (Exp.AddComponent(GUI.NewBar(17))).(*GUI.Bar)

	money := engine.NewGameObject("money")
	money.Transform().SetWorldPositionf(100, 500)
	money.Transform().SetScalef(20, 20)
	money.Transform().SetDepth(2)
	Ch.Money = money.AddComponent(components.NewUIText(Fonts.ArialFont2, "0")).(*components.UIText)
	Ch.Money.SetAlign(engine.AlignLeft)

	level := engine.NewGameObject("level")
	level.Transform().SetWorldPositionf(50, 500)
	level.Transform().SetScalef(20, 20)
	level.Transform().SetDepth(2)
	Ch.Level = level.AddComponent(components.NewUIText(Fonts.ArialFont2, "1")).(*components.UIText)
	Ch.Level.SetAlign(engine.AlignLeft)
}
func NewPlayer() *Player {
	return &Player{engine.NewComponent(), 0, 0, 1, 50, 0, 100, 100, 100, 100, 100, 50, stand_height, 60, 7000, false, false, false, 1, true, false, true, nil, nil, nil, engine.StartCoroutine(func() {}), nil}
}
func (pl *Player) AddMoney(value int) {
	pl.money += value
	Ch.Money.SetString(strconv.Itoa(pl.money))
}
func (pl *Player) AddExp(value float32) {
	pl.Exp += value
	for pl.Exp > pl.MaxExp {
		pl.Exp = pl.Exp - pl.MaxExp
		pl.level++
		Ch.Level.SetString(strconv.Itoa(pl.level))
	}
	Ch.Exp.SetValue(pl.Exp, pl.MaxHp)
}
func (pl *Player) Start() {
	pl.GameObject().Physics.Body.SetMoment(engine.Inf)
}
func (pl *Player) AddHp(val float32) {
	pl.Hp = float32(math.Min(float64(pl.MaxHp), float64(pl.Hp+val)))
	Ch.Hp.SetValue(pl.Hp, pl.MaxHp)
}
func (pl *Player) OnCollisionEnter(arbiter engine.Arbiter) bool {
	if arbiter.GameObjectB() != nil {
		if arbiter.GameObjectB().Tag == "lader" {
			pl.pLader = arbiter.GameObjectB()
			if pl.GameObject().Sprite.CurrentAnimation() != "player_climb" {
				pl.GameObject().Sprite.SetAnimation("player_climb")
			}
			pl.GameObject().Physics.Body.IgnoreGravity = true
			pl.OnGround = false
		}
		if arbiter.GameObjectB().Tag == "splinter" {
			pl.pSplint = arbiter.GameObjectB()
			engine.StartCoroutine(func() {
				for pl.pSplint != nil {
					if pl.Hitable {
						pl.Hit(15)
					}
					engine.CoYieldCoroutine(pl.Hitted)

				}

			})
		}
	}
	return true
}
func (pl *Player) Hit(dmg int) {
	if pl.Hitted.State == engine.Ended {
		pl.Hitted = engine.StartCoroutine(func() {
			pl.hit = true

			GUI.NewUpTextObj(strconv.Itoa(dmg), pl.Transform(), 20)
			pl.LastFloor = nil
			pl.Hitable = false
			pl.Attack = false
			pl.able = false
			pl.GameObject().Physics.Body.AddForce(0, pl.jumpPower)
			pl.SubLife(float32(dmg))
			engine.CoSleep(3)
			pl.hit = false
			pl.able = true
			pl.GameObject().Sprite.SetAnimation("player_stand")
			engine.CoSleep(2)
			pl.Hitable = true
		})
	}
}
func (pl *Player) SubLife(hp float32) {
	pl.Hp -= hp
	if pl.Hp < 0 {
		pl.Hp = 0
	}
	Ch.Hp.SetValue(pl.Hp, pl.MaxHp)

}

func (pl *Player) OnCollisionPostSolve(arbiter engine.Arbiter) {
	if arbiter.GameObjectB() == nil {
		return
	}
	if arbiter.GameObjectB().Tag != "lader" && arbiter.GameObjectA().Tag != "lader" {
		count := 0
		for _, con := range arbiter.Contacts {
			if arbiter.Normal(con).Y < -0.9 {
				count++

			}
		}
		if count >= 1 {
			if pl.GameObject().Sprite.CurrentAnimation() == "player_jump" {
				pl.GameObject().Sprite.SetAnimation("player_stand")
			}
			pl.LastFloor = arbiter.GameObjectB()

			pl.OnGround = true
		}
	}

}
func (pl *Player) FixedUpdate() {
	pl.OnGround = false
	pl.LastFloor = nil
}

func (pl *Player) OnCollisionExit(arbiter engine.Arbiter) {
	if arbiter.GameObjectB() == pl.pSplint {
		pl.pSplint = nil
	} else if arbiter.GameObjectB() == pl.pLader {
		pl.pLader = nil
		pl.GameObject().Physics.Body.IgnoreGravity = false
		pl.GameObject().Sprite.SetAnimation("player_stand")
	}
}
func (pl *Player) Update() {
	//Test
	if pl.OnGround == false {
		pl.frames++
	} else {
		pl.frames = 0
	}
	// if input.KeyPress(input.KeyEsc) {
	// 	engine.LoadScene(MenuSceneG)
	// }
	if input.KeyDown(input.KeyEsc) {
		pl.MenuScene()
	}
	ph := pl.GameObject().Physics.Body
	pl.GameObject().Sprite.SetAlign(engine.AlignTopCenter)
	if float32(math.Abs(float64(ph.Velocity().X))) > 3 {
		if pl.GameObject().Sprite.CurrentAnimation() == "player_stand" {
			pl.GameObject().Sprite.SetAnimation("player_walk")
		}
	} else if !pl.Attack && pl.pLader == nil {
		pl.GameObject().Sprite.SetAnimation("player_stand")
	}
	if pl.able {
		if input.KeyDown(input.Key_Right) {
			ph.AddForce(pl.speed, 0)
			pl.right = 1
			pl.GameObject().Transform().SetScalef(pl.width, pl.height)

		} else if input.KeyDown(input.Key_Left) {
			ph.AddForce(-pl.speed, 0)
			pl.right = -1
			pl.GameObject().Transform().SetScalef(-pl.width, pl.height)
		}
		if input.KeyPress(input.KeyLctrl) && !pl.Attack {
			pl.Attack = true

			pl.GameObject().Sprite.SetAnimation("player_attack")
			pl.GameObject().Sprite.AnimationEndCallback = func(sprite *engine.Sprite) {
				pl.Attack = false
				pl.GameObject().Sprite.SetAnimation("player_stand")
				pl.GameObject().Sprite.AnimationEndCallback = nil
			}

		}
		if input.KeyPress(input.Key_Up) && pl.OnGround {
			pl.GameObject().Physics.Body.AddForce(0, pl.jumpPower)
			pl.OnGround = false
		}
		if input.KeyDown(input.Key_Up) {
			if pl.pLader != nil {
				ph.AddVelocity(0, 1)
			}
		}
		if input.KeyUp(input.Key_Down) {
			if pl.GameObject().Sprite.CurrentAnimation() == "player_bend" {
				pl.GameObject().Sprite.SetAnimation("player_stand")
			}
			pl.GameObject().Physics.Shape.SetFriction(0.7)
			pl.height = stand_height
			pl.GameObject().Transform().SetScalef(pl.width*pl.right, stand_height)

		} else if input.KeyDown(input.Key_Down) {
			if pl.pLader != nil {
				ph.AddVelocity(0, -1)
			} else {
				pl.GameObject().Sprite.SetAnimation("player_bend")
				pl.GameObject().Physics.Shape.SetFriction(1.2)
				pl.height = bend_height
				pl.GameObject().Transform().SetScalef(pl.width*pl.right, bend_height)
			}
		}
	}

	if pl.hit {
		pl.GameObject().Sprite.SetAnimation("player_hit")
	}
	if ((pl.GameObject().Sprite.CurrentAnimation() == "player_stand") || (pl.GameObject().Sprite.CurrentAnimation() == "player_walk")) && !pl.OnGround && pl.frames > 15 {
		pl.GameObject().Sprite.SetAnimation("player_jump")
	}
	if pl.GameObject().Sprite.CurrentAnimation() != "player_climb" && pl.pLader != nil {
		pl.GameObject().Sprite.SetAnimation("player_climb")
	}

}
