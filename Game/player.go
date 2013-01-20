package Game

import (
	"github.com/vova616/GarageEngine/Engine"
	"github.com/vova616/GarageEngine/Engine/Input"
)

type Player struct {
	Engine.BaseComponent
	width    float32
	height   float32
	speed    float32
	Atack    bool
	OnGround bool
}

func NewPlayer() *Player {
	return &Player{Engine.NewComponent(), 50, 100, 60, false, false}
}
func (pl *Player) Start() {
}
func (pl *Player) Update() {

	pl.GameObject().Physics.Body.AddForce(0, -100)
	ph := pl.GameObject().Physics.Body
	if Input.KeyDown(Input.Key_Right) {
		ph.AddForce(pl.speed, 0)
		pl.GameObject().Transform().SetScalef(pl.width, pl.height)
		//pl.GameObject().Sprite.BindAnimations(player_walk)
	} else if Input.KeyDown(Input.Key_Left) {
		ph.AddForce(-pl.speed, 0)
		pl.GameObject().Transform().SetScalef(-pl.width, pl.height)
		//pl.GameObject().Sprite.BindAnimations(player_walk)
	} else {
		//pl.GameObject().Sprite.BindAnimations(player_stand)
	}
	if Input.KeyPress(Input.KeyLctrl) {
		pl.Atack = true
		/*
			pl.GameObject().Sprite.BindAnimations(player_atack)
			pl.GameObject().Sprite.AnimationEndCallback = func(sprite *Engine.Sprite) {
				pl.Atack = false
				pl.GameObject().Sprite.BindAnimations(player_stand)
			}
		*/
	}
	if Input.KeyPress(Input.Key_Up) /*&& OnGround*/ {
		pl.GameObject().Physics.Body.AddForce(0, 5000)
	}
	/*if(!OnGround) {
		pl.GameObject().Sprite.BindAnimations(player_jump)
	}*/

}
