package Objects

import (
	"fmt"
	"github.com/vova616/garageEngine/engine"
	"math/rand"
)

type Bird struct {
	engine.BaseComponent
	dir   bool
	speed float32
}

func NewBird(dir bool, speed float32) *Bird {
	return &Bird{engine.NewComponent(), dir, speed}
}
func (b *Bird) Start() {
	cp := engine.GetScene().SceneBase().Camera.Transform().WorldPosition()
	sc := b.Transform().Scale()
	if b.dir {
		b.Transform().SetPosition(cp.Add(engine.Vector{-100, float32(rand.Intn(300)) + 500, 0}))
		sc.X *= -1
	} else {
		b.Transform().SetPosition(cp.Add(engine.Vector{1500, float32(rand.Intn(300)) + 500, 0}))
	}

	b.Transform().SetScale(sc)
}
func (b *Bird) Update() {
	p := b.Transform().WorldPosition()
	if b.dir {
		p.X += b.speed
	} else {
		p.X -= b.speed
	}
	b.Transform().SetWorldPosition(p)
	cp := engine.GetScene().SceneBase().Camera.Transform().WorldPosition()
	if cp.Distance(p) > 4000 {
		b.GameObject().Destroy()
	}
	fmt.Println(b.Transform().Position())
}
