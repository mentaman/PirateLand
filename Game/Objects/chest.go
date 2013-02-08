package Objects

import (
	"github.com/vova616/garageEngine/engine"
	"github.com/vova616/garageEngine/engine/input"

	"github.com/mentaman/PirateLand/Game/Player"
	"math/rand"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

const (
	Type_money  = 1
	Type_prices = 2
)

type Chest struct {
	engine.BaseComponent
	priceType int
	playerIn  bool
	done      bool
}

func NewChest(ty int) *Chest {
	return &Chest{engine.NewComponent(), ty, false, false}
}

func (c *Chest) Update() {
	if !c.done && c.playerIn && input.KeyDown('E') {
		c.done = true
		switch c.priceType {
		case Type_money:
			Player.PlComp.AddMoney(rand.Int()%1000 + 1)
			c.GameObject().Sprite.AnimationSpeed = 5

			c.GameObject().Sprite.AnimationEndCallback = func(sprite *engine.Sprite) {
				c.GameObject().Destroy()
			}
		case Type_prices:
			r := rand.Int()%6 + 2
			for i := 0; i < r; i++ {
				it := RandomItem().Clone()

				it.Transform().SetPosition(c.Transform().Position())
				it.Transform().SetParent(c.Transform().Parent())
				it.ComponentTypeOfi((*Item).TypeOf(nil)).(*Item).Pop()
			}
			c.GameObject().Sprite.AnimationSpeed = 5
			c.GameObject().Sprite.AnimationEndCallback = func(sprite *engine.Sprite) {
				c.GameObject().Destroy()
			}
		default:
			c.done = false
		}

	}

}
func (c *Chest) Start() {
	if c.priceType == -1 {
		c.priceType = rand.Int()%2 + 1
	}
}
func (c *Chest) OnCollisionEnter(arbiter engine.Arbiter) bool {

	if arbiter.GameObjectB().Tag == "player" {
		c.playerIn = true
	}
	return true
}
func (c *Chest) OnCollisionExit(arbiter engine.Arbiter) {
	if arbiter.GameObjectB() != nil {
		if arbiter.GameObjectB().Tag == "player" {
			c.playerIn = false
		}
	}
}
