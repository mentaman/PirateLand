package Objects

import (
	"github.com/vova616/garageEngine/engine"
)

var (
	Floor    *engine.GameObject
	Splinter *engine.GameObject
	Lader    *engine.GameObject
	Box      *engine.GameObject
	ChestO   *engine.GameObject
	BirdO    *engine.GameObject

	BirdControll *engine.Coroutine

	Tileset      *engine.ManagedAtlas
	ObjectsAtlas *engine.ManagedAtlas
)

const (
	Spr_lader    = 1
	Spr_splinter = 2
	Spr_box      = 3
	Spr_bird     = 4
)

func createFloor() {

	uvs, ind := engine.AnimatedGroupUVs(Tileset, "ground")
	Floor = engine.NewGameObject("floor")
	Floor.AddComponent(engine.NewSprite3(Tileset.Texture, uvs))
	Floor.Sprite.BindAnimations(ind)
	Floor.Transform().SetDepth(-1)
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
func createBird() {
	BirdO = engine.NewGameObject("bird")
	BirdO.AddComponent(engine.NewSprite2(ObjectsAtlas.Texture, engine.IndexUV(ObjectsAtlas, Spr_bird)))
	BirdO.Transform().SetScalef(50, 50)
	if BirdControll != nil {
		BirdControll.State = engine.Ended
		BirdControll = nil
	}

}
func createChest() {
	uvs, ind := engine.AnimatedGroupUVs(ObjectsAtlas, "chest")
	ChestO = engine.NewGameObject("chest")
	ChestO.AddComponent(engine.NewSprite3(ObjectsAtlas.Texture, uvs))
	ChestO.Sprite.BindAnimations(ind)
	ChestO.Transform().SetWorldScalef(70, 70)
	ChestO.AddComponent(engine.NewPhysics(false, 1, 1))
	ChestO.Physics.Shape.IsSensor = true
	ChestO.Sprite.AnimationSpeed = 0
	ChestO.Physics.Body.IgnoreGravity = true
	ChestO.AddComponent(NewChest(-1))
}
func CreateObjects() {
	createFloor()
	createSplinter()
	createBox()
	createLader()
	initItems()
	createChest()
	createBird()
}
