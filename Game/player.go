package Game

import (
	"github.com/vova616/GarageEngine/Engine"
	"github.com/vova616/GarageEngine/Engine/Input"
	"math"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

type Player struct {
	Engine.BaseComponent
	width     float32
	height    float32
	speed     float32
	jumpPower float32
	Attack    bool
	OnGround  bool
	LastFloor *Engine.GameObject
}

func NewPlayer() *Player {
	return &Player{Engine.NewComponent(), 50, 100, 60, 5000, false, false, nil}
}
func (pl *Player) Start() {
}
func (pl *Player) OnCollisionEnter(arbiter *Engine.Arbiter) bool {
	count := 0
	for _, con := range arbiter.Contacts {
		if con.Normal().Y > 0.9 {
			count++
		}
	}
	if count >= 2 {
		pl.OnGround = true
		pl.GameObject().Sprite.SetAnimation("player_stand")
		pl.LastFloor = arbiter.GameObjectB()
	}

	return true
}
func (pl *Player) OnCollisionExit(arbiter *Engine.Arbiter) {
	if arbiter.GameObjectB() == pl.LastFloor {

		count := 0
		for _, con := range arbiter.Contacts {
			if con.Normal().Y > 0.9 {
				count++
			}
		}
		if count < 2 {
			pl.OnGround = false
			pl.GameObject().Sprite.SetAnimation("player_stand")
		}
	}
}
func (pl *Player) Update() {
	//Test

	ph := pl.GameObject().Physics.Body
	if float32(math.Abs(float64(ph.Velocity().X))) > 3 {
		if pl.GameObject().Sprite.CurrentAnimation() == "player_stand" {
			pl.GameObject().Sprite.SetAnimation("player_walk")
		}
	} else if !pl.Attack {
		pl.GameObject().Sprite.SetAnimation("player_stand")
	}
	ph.SetAngularVelocity(0)
	if Input.KeyDown(Input.Key_Right) {
		ph.AddForce(pl.speed, 0)
		pl.GameObject().Transform().SetScalef(pl.width, pl.height)

	} else if Input.KeyDown(Input.Key_Left) {
		ph.AddForce(-pl.speed, 0)
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
	if !pl.OnGround {
		pl.GameObject().Sprite.SetAnimation("player_jump")
	}

}
