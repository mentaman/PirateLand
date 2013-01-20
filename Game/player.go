package Game

import (
	"github.com/vova616/GarageEngine/Engine"
	"github.com/vova616/GarageEngine/Engine/Input"
)

type Player struct {
	Engine.BaseComponent
	width  float32
	height float32
	speed  float32
}

func NewPlayer() *Player {
	return &Player{Engine.NewComponent(), 50, 100, 60}
}
func (pl *Player) Start() {
}
func (pl *Player) Update() {

	pl.GameObject().Physics.Body.AddForce(0, -20)
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

}
