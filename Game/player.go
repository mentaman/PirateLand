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
<<<<<<< HEAD
	Attack   bool
=======
	Attack    bool
>>>>>>> ab1c590a098d2c66cc195ddccd0c7e91fa476e44
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
<<<<<<< HEAD
	}
=======
	} 
>>>>>>> ab1c590a098d2c66cc195ddccd0c7e91fa476e44
	pl.GameObject().Physics.Body.AddForce(0, -100)
	ph := pl.GameObject().Physics.Body
	ph.SetAngularVelocity(0)
	if Input.KeyDown(Input.Key_Right) {
		ph.AddForce(pl.speed, 0)
		pl.GameObject().Transform().SetScalef(pl.width, pl.height)
	} else if Input.KeyDown(Input.Key_Left) {
		ph.AddForce(-pl.speed, 0)
		pl.GameObject().Transform().SetScalef(-pl.width, pl.height)
<<<<<<< HEAD
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

=======
	} else if !pl.Attack{
		pl.GameObject().Sprite.SetAnimation("player_stand")
	}
	if Input.KeyPress(Input.KeyLctrl) {
		pl.Atack = true
	
		pl.GameObject().Sprite.SetAnimation("player_atack")
		pl.GameObject().Sprite.AnimationEndCallback = func(sprite *Engine.Sprite) {
			pl.Atack = false
			pl.GameObject().Sprite.SetAnimation("player_stand")
		}
		
>>>>>>> ab1c590a098d2c66cc195ddccd0c7e91fa476e44
	}
	if Input.KeyPress(Input.Key_Up) && pl.OnGround {
		pl.GameObject().Physics.Body.AddForce(0, 5000)
		pl.OnGround = false
	}
<<<<<<< HEAD
	if !pl.OnGround {
=======
	if(!pl.OnGround) {
>>>>>>> ab1c590a098d2c66cc195ddccd0c7e91fa476e44
		pl.GameObject().Sprite.SetAnimation("player_jump")
	}

}
