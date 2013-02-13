package Objects

import (
	"github.com/mentaman/PirateLand/Game/GUI"
	"github.com/mentaman/PirateLand/Game/Player"
	"github.com/vova616/garageEngine/engine"
	"github.com/vova616/garageEngine/engine/components/tween"
	"math/rand"
	"strconv"
	"time"

//	"github.com/vova616/chipmunk/vect"

//	"github.com/vova616/chipmunk"
)

var (
	Atlas *engine.ManagedAtlas
)
var Items = struct {
	Spot    *engine.GameObject
	BigSpot *engine.GameObject
	Coin    *engine.GameObject
	Coin10  *engine.GameObject
	Diamond *engine.GameObject
}{engine.NewGameObject("spot"), engine.NewGameObject("bigspot"), engine.NewGameObject("coin"), engine.NewGameObject("coin10"), engine.NewGameObject("diamond")}

const (
	Spr_coin    = 1
	Spr_coin10  = 2
	Spr_diamond = 3
	Spr_spot    = 4
	Spr_bigspot = 4
)

type Item struct {
	engine.BaseComponent
	coll     func(*engine.GameObject)
	takeable bool
}

func NewItem(coll func(*engine.GameObject)) *Item {
	return &Item{engine.NewComponent(), coll, false}
}
func (this *Item) TypeOf() interface{} {
	return this
}
func (s *Item) Pop() {
	s.GameObject().Physics.Body.AddForce(float32(rand.Int()%30-15), float32(rand.Int()%3000+4000))
}
func (s *Item) Start() {
	engine.StartCoroutine(func() {
		engine.CoSleep(2)
		s.takeable = true
		engine.CoSleep(6)
		tween.Create(&tween.Tween{Target: s.GameObject(), From: []float32{1}, To: []float32{0.2},
			Algo: tween.Linear, Type: tween.Color, Time: time.Second * 4, Loop: tween.None, Format: "a", EndCallback: func() {
				if s.GameObject() != nil {
					s.GameObject().Destroy()
				}
			}})

	})

}
func (s *Item) OnCollisionEnter(arbiter engine.Arbiter) bool {
	if arbiter.GameObjectB() != nil {
		if arbiter.GameObjectB().Tag == "player" {
			if s.takeable {
				s.coll(s.GameObject())
				s.GameObject().Destroy()
				s.takeable = false
			}
		}
	}
	return true
}
func RandomItem() *engine.GameObject {
	i := rand.Int() % 5
	switch i {
	case 0:
		return Items.Spot
	case 1:
		return Items.Coin
	case 2:
		return Items.Coin10
	case 3:
		return Items.Diamond
	case 4:
		return Items.BigSpot
	}

	return Items.Spot
}
func initItems() {
	Items.Spot.AddComponent(engine.NewSprite2(Atlas.Texture, engine.IndexUV(Atlas, Spr_spot)))
	Items.Spot.AddComponent(engine.NewPhysics(false, 1, 1))
	Items.Spot.Transform().SetScalef(30, 30)
	Items.Spot.AddComponent(NewItem(func(so *engine.GameObject) {
		r := rand.Int()%10 + 5
		Player.PlComp.AddHp(float32(r))

		GUI.NewUpTextObj(strconv.Itoa(r), so.Transform(), 20)
	}))

	Items.BigSpot.AddComponent(engine.NewSprite2(Atlas.Texture, engine.IndexUV(Atlas, Spr_bigspot)))
	Items.BigSpot.AddComponent(engine.NewPhysics(false, 1, 1))
	Items.BigSpot.Transform().SetScalef(30, 30)
	Items.BigSpot.AddComponent(NewItem(func(so *engine.GameObject) {
		r := rand.Int()%15 + 15
		Player.PlComp.AddHp(float32(r))

		GUI.NewUpTextObj(strconv.Itoa(r), so.Transform(), 20)
	}))

	Items.Coin.AddComponent(engine.NewSprite2(Atlas.Texture, engine.IndexUV(Atlas, Spr_coin)))
	Items.Coin.AddComponent(engine.NewPhysics(false, 1, 1))
	Items.Coin.Transform().SetScalef(30, 30)
	Items.Coin.AddComponent(NewItem(func(so *engine.GameObject) {
		Player.PlComp.AddMoney(1)

		GUI.NewUpTextObj("1", so.Transform(), 20)
	}))

	Items.Coin10.AddComponent(engine.NewSprite2(Atlas.Texture, engine.IndexUV(Atlas, Spr_coin10)))
	Items.Coin10.AddComponent(engine.NewPhysics(false, 1, 1))
	Items.Coin10.Transform().SetScalef(30, 30)
	Items.Coin10.AddComponent(NewItem(func(so *engine.GameObject) {
		Player.PlComp.AddMoney(10)
		GUI.NewUpTextObj("10", so.Transform(), 20)
	}))

	Items.Diamond.AddComponent(engine.NewSprite2(Atlas.Texture, engine.IndexUV(Atlas, Spr_diamond)))
	Items.Diamond.AddComponent(engine.NewPhysics(false, 1, 1))
	Items.Diamond.Transform().SetScalef(30, 30)
	Items.Diamond.AddComponent(NewItem(func(so *engine.GameObject) {
		r := (rand.Intn(50) + 20)
		Player.PlComp.AddMoney(int(r))

		GUI.NewUpTextObj(strconv.Itoa(r), so.Transform(), 20)
	}))
}
