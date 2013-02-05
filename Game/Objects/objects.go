package Objects

import (
	"github.com/mentaman/PirateLand/Game/Player"
	"github.com/vova616/garageEngine/engine"
	"math/rand"
)

var (
	Floor    *engine.GameObject
	Splinter *engine.GameObject
	Lader    *engine.GameObject
	Box      *engine.GameObject

	Tileset      *engine.ManagedAtlas
	ObjectsAtlas *engine.ManagedAtlas
)

const (
	Spr_lader    = 1
	Spr_splinter = 2
	Spr_box      = 3
)

func createFloor() {

	uvs, ind := engine.AnimatedGroupUVs(Tileset, "ground")
	Floor = engine.NewGameObject("floor")
	Floor.AddComponent(engine.NewSprite3(Tileset.Texture, uvs))
	Floor.Sprite.BindAnimations(ind)
	Floor.Sprite.AnimationSpeed = 0
	Floor.Transform().SetWorldScalef(100, 100)
	Floor.AddComponent(engine.NewPhysics(true, 1, 1))
}
func createSplinter() {
	Splinter = engine.NewGameObject("Splinter")
	Splinter.AddComponent(engine.NewSprite2(ObjectsAtlas.Texture, engine.IndexUV(ObjectsAtlas, Spr_splinter)))
	Splinter.Transform().SetWorldScalef(100, 30)
	Splinter.AddComponent(engine.NewPhysics(true, 1, 1))
	Splinter.Physics.Shape.IsSensor = true
	Splinter.Tag = "splinter"
}
func createBox() {
	Box = engine.NewGameObject("box")
	Box.AddComponent(engine.NewSprite2(ObjectsAtlas.Texture, engine.IndexUV(ObjectsAtlas, Spr_box)))
	Box.Transform().SetWorldScalef(40, 40)
	Box.AddComponent(engine.NewPhysics(false, 1, 1))
	Box.Physics.Body.SetMass(1.5)
	Box.Physics.Shape.SetFriction(0.3)
}
func createLader() {
	Lader = engine.NewGameObject("lader")
	Lader.AddComponent(engine.NewSprite2(ObjectsAtlas.Texture, engine.IndexUV(ObjectsAtlas, Spr_lader)))
	Lader.Transform().SetWorldScalef(60, 100)
	Lader.AddComponent(engine.NewPhysics(true, 1, 1))
	Lader.Physics.Shape.IsSensor = true
	Lader.Physics.Shape.SetFriction(2)
	Lader.Tag = "lader"
}
func createSpot() {
	Spot = engine.NewGameObject("spot")
	Spot.AddComponent(engine.NewSprite2(Atlas.Texture, engine.IndexUV(Atlas, Spr_spot)))
	Spot.AddComponent(engine.NewPhysics(false, 1, 1))
	Spot.Transform().SetScalef(30, 30)
	Spot.AddComponent(NewItem(func(so *engine.GameObject) {
		Player.PlComp.AddHp(float32(rand.Int()%10 + 5))
		so.Destroy()
	}))
}
func CreateObjects() {
	createFloor()
	createSplinter()
	createBox()
	createLader()
	createSpot()
}
