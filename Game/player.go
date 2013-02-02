package Game

import (
	"github.com/vova616/GarageEngine/Engine"
	"github.com/vova616/GarageEngine/Engine/Input"
	"math"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

var (
	plComp *Player
)

type Player struct {
	Engine.BaseComponent
	frames    int
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
	hitable   bool
	pLader    *Engine.GameObject
	pSplint   *Engine.GameObject
	LastFloor *Engine.GameObject
	Hitted    *Engine.Coroutine
}

const stand_height = 100
const bend_height = 70

func NewPlayer() *Player {
	return &Player{Engine.NewComponent(), 0, 100, 100, 50, stand_height, 60, 7000, false, false, false, 1, true, false, true, nil, nil, nil, nil}
}
func (pl *Player) Start() {
	plComp = pl
	pl.GameObject().Physics.Body.SetMoment(Engine.Inf)
}
func (pl *Player) OnCollisionEnter(arbiter Engine.Arbiter) bool {

	if arbiter.GameObjectB().Tag == "lader" {
		pl.pLader = arbiter.GameObjectB()
		if pl.GameObject().Sprite.CurrentAnimation() != "player_climb" {
			pl.GameObject().Sprite.SetAnimation("player_climb")
		}
		pl.GameObject().Physics.Body.IgnoreGravity = true
		pl.OnGround = false
	}
	if pl.hitable && arbiter.GameObjectB().Tag == "splinter" {
		pl.pSplint = arbiter.GameObjectB()
		Engine.StartCoroutine(func() {
			for pl.pSplint != nil {
				pl.Hit()
				Engine.CoYieldCoroutine(pl.Hitted)
			}
		})
	}
	return true
}
func (pl *Player) Hit() {

	pl.Hitted = Engine.StartCoroutine(func() {
		pl.hit = true
		pl.LastFloor = nil
		pl.hitable = false
		pl.Attack = false
		pl.able = false
		pl.GameObject().Physics.Body.AddForce(0, pl.jumpPower)
		pl.SubLife(5)
		Engine.CoSleep(3)
		pl.hit = false
		pl.able = true
		pl.GameObject().Sprite.SetAnimation("player_stand")
		Engine.CoSleep(2)
		pl.hitable = true
	})
}
func (pl *Player) SubLife(hp float32) {
	pl.Hp -= hp
	if pl.Hp < 0 {
		pl.Hp = 0
	}
	ch.Hp.SetValue(pl.Hp / pl.MaxHp)

}

func (pl *Player) OnCollisionPostSolve(arbiter Engine.Arbiter) {
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

func (pl *Player) OnCollisionExit(arbiter Engine.Arbiter) {
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
	if Input.KeyPress(Input.KeyEsc) {
		Engine.LoadScene(MenuSceneG)
	}
	ph := pl.GameObject().Physics.Body
	pl.GameObject().Sprite.SetAlign(Engine.AlignTopCenter)
	if float32(math.Abs(float64(ph.Velocity().X))) > 3 {
		if pl.GameObject().Sprite.CurrentAnimation() == "player_stand" {
			pl.GameObject().Sprite.SetAnimation("player_walk")
		}
	} else if !pl.Attack && pl.pLader == nil {
		pl.GameObject().Sprite.SetAnimation("player_stand")
	}
	if pl.able {
		if Input.KeyDown(Input.Key_Right) {
			ph.AddForce(pl.speed, 0)
			pl.right = 1
			pl.GameObject().Transform().SetScalef(pl.width, pl.height)

		} else if Input.KeyDown(Input.Key_Left) {
			ph.AddForce(-pl.speed, 0)
			pl.right = -1
			pl.GameObject().Transform().SetScalef(-pl.width, pl.height)
		}
		if Input.KeyPress(Input.KeyLctrl) {
			pl.Attack = true

			pl.GameObject().Sprite.SetAnimation("player_attack")
			pl.GameObject().Sprite.AnimationEndCallback = func(sprite *Engine.Sprite) {
				pl.Attack = false
				pl.GameObject().Sprite.SetAnimation("player_stand")
			}

		}
		if Input.KeyPress(Input.Key_Up) && pl.OnGround {
			pl.GameObject().Physics.Body.AddForce(0, pl.jumpPower)
			pl.OnGround = false
		}
		if Input.KeyDown(Input.Key_Up) {
			if pl.pLader != nil {
				ph.AddVelocity(0, 1)
			}
		}
		if Input.KeyUp(Input.Key_Down) {
			if pl.GameObject().Sprite.CurrentAnimation() == "player_bend" {
				pl.GameObject().Sprite.SetAnimation("player_stand")
			}
			pl.height = stand_height
			pl.GameObject().Transform().SetScalef(pl.width*pl.right, stand_height)

		} else if Input.KeyDown(Input.Key_Down) {
			if pl.pLader != nil {
				ph.AddVelocity(0, -1)
			} else {
				pl.GameObject().Sprite.SetAnimation("player_bend")
				pl.height = bend_height
				pl.GameObject().Transform().SetScalef(pl.width*pl.right, bend_height)
			}
		}
	}

	if pl.hit {
		pl.GameObject().Sprite.SetAnimation("player_hit")
	}
	if (pl.GameObject().Sprite.CurrentAnimation() == "player_stand" || pl.GameObject().Sprite.CurrentAnimation() == "player_walk") && !pl.OnGround && pl.frames > 15 {
		pl.GameObject().Sprite.SetAnimation("player_jump")
	}
	if pl.GameObject().Sprite.CurrentAnimation() != "player_climb" && pl.pLader != nil {
		pl.GameObject().Sprite.SetAnimation("player_climb")
	}
}
