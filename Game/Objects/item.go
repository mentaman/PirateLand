package Objects

import (
	"github.com/vova616/garageEngine/engine"
	"math/rand"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

var (
	Atlas *engine.ManagedAtlas
	Spot  *engine.GameObject
)

const (
	Spr_coin    = 1
	Spr_coin10  = 2
	Spr_diamond = 3
	Spr_spot    = 4
)

type Item struct {
	engine.BaseComponent
	coll func(*engine.GameObject)
}

func NewItem(coll func(*engine.GameObject)) *Item {
	return &Item{engine.NewComponent(), coll}
}
func (s *Item) Pop() {
	s.GameObject().Physics.Body.AddForce(float32(rand.Int()%30-15), float32(rand.Int()%3000+4000))
}
func (s *Item) OnCollisionEnter(arbiter engine.Arbiter) bool {
	if arbiter.GameObjectB().Tag == "player" {
		s.coll(s.GameObject())
	}
	return true
}
