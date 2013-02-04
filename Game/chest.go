package Game

import (
	"github.com/vova616/garageEngine/engine"
	"github.com/vova616/garageEngine/engine/input"
	"math/rand"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

type typeGift byte

const (
	money  = typeGift(1)
	prices = typeGift(2)
)

type Chest struct {
	engine.BaseComponent
	priceType typeGift
	playerIn  bool
	done      bool
}

func NewChest(ty typeGift) *Chest {
	return &Chest{engine.NewComponent(), ty, false, false}
}

func (c *Chest) Update() {
	if !c.done && c.playerIn && input.KeyDown('E') {
		c.done = true
		switch c.priceType {
		case money:
			plComp.AddMoney(rand.Int()%1000 + 1)
			c.GameObject().Sprite.AnimationSpeed = 5

			c.GameObject().Sprite.AnimationEndCallback = func(sprite *engine.Sprite) {
				c.GameObject().Destroy()
			}
		default:
			c.done = false
		}

	}

}
func (c *Chest) OnCollisionEnter(arbiter engine.Arbiter) bool {

	if arbiter.GameObjectB().Tag == "player" {
		c.playerIn = true
	}
	return true
}
func (c *Chest) OnCollisionExit(arbiter engine.Arbiter) {
	if arbiter.GameObjectB().Tag == "player" {
		c.playerIn = false
	}
}
