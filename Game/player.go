package Game

import (
	"github.com/vova616/GarageEngine/Engine"
	"github.com/vova616/GarageEngine/Engine/Input"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

type Player struct {
	Engine.BaseComponent
	width    float32
	height   float32
	speed    float32
	Attack   bool
	OnGround bool
}

func NewPlayer() *Player {
	return &Player{Engine.NewComponent(), 50, 100, 60, false, false}
}
func (pl *Player) Start() {
}
func (pl *Player) OnCollisionEnter(arbiter *Engine.Arbiter) bool {
	for _, con := range arbiter.Contacts {
		if con.Normal().Y == 1 {
			pl.OnGround = true
		}
	}
	return true
}

func (pl *Player) Update() {
	if Input.KeyPress(Input.Key_Right) || Input.KeyPress(Input.Key_Right) {
		pl.GameObject().Sprite.SetAnimation("player_walk")
	}
	pl.GameObject().Physics.Body.AddForce(0, -100)
	ph := pl.GameObject().Physics.Body
	ph.SetAngularVelocity(0)
	if Input.KeyDown(Input.Key_Right) {
		ph.AddForce(pl.speed, 0)
		pl.GameObject().Transform().SetScalef(pl.width, pl.height)
	} else if Input.KeyDown(Input.Key_Left) {
		ph.AddForce(-pl.speed, 0)
		pl.GameObject().Transform().SetScalef(-pl.width, pl.height)
	} else if !pl.Attack {
		pl.GameObject().Sprite.SetAnimation("player_stand")
	}
	if Input.KeyPress(Input.KeyLctrl) {
		pl.Atatck = true

		pl.GameObject().Sprite.SetAnimation("player_attack")
		pl.GameObject().Sprite.AnimationEndCallback = func(sprite *Engine.Sprite) {
			pl.Attack = false
			pl.GameObject().Sprite.SetAnimation("player_stand")
		}

	}
	if Input.KeyPress(Input.Key_Up) && pl.OnGround {
		pl.GameObject().Physics.Body.AddForce(0, 5000)
		pl.OnGround = false
	}
	if !pl.OnGround {
		pl.GameObject().Sprite.SetAnimation("player_jump")
	}

}
