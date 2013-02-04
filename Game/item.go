package Game

import (
	"github.com/vova616/garageEngine/engine"
	"math/rand"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

type Item struct {
	engine.BaseComponent
	coll func()
}

func NewItem(coll func()) *Item {
	return &Item{engine.NewComponent(), coll}
}
func (s *Item) Pop() {
	s.GameObject().Physics.Body.AddForce(float32(rand.Int()%30-15), float32(rand.Int()%3000+4000))
}
func (s *Item) OnCollisionEnter(arbiter engine.Arbiter) bool {
	if arbiter.GameObjectB().Tag == "player" {
		s.coll()
	}
	return true
}
