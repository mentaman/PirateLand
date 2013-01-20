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
	return &Player{Engine.NewComponent(), 50, 100, 30}
}
func (pl *Player) Start() {
}
func (pl *Player) Update() {

	ph := pl.GameObject().Physics
	if Input.KeyDown(Input.Key_Right) {
		ph.Body.SetVelocity(pl.speed, float32(ph.Body.Velocity().Y))
		pl.GameObject().Transform().SetScalef(pl.width, pl.height)
		//pl.GameObject().Sprite.Bind(player_walk)
	} else if Input.KeyDown(Input.Key_Left) {
		ph.Body.SetVelocity(-pl.speed, float32(ph.Body.Velocity().Y))
		pl.GameObject().Transform().SetScalef(-pl.width, pl.height)
		//pl.GameObject().Sprite.Bind(player_walk)
	} else {
		ph.Body.SetVelocity(0, 0)
		//pl.GameObject().Sprite.Bind(player_stand)
	}

}
