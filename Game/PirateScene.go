package Game

import (
	"fmt"
	"github.com/vova616/GarageEngine/Engine"
	"github.com/vova616/GarageEngine/Engine/Components"
	"math/rand"
	"time"
)

var (
	atlas            *Engine.ManagedAtlas
	plAtlas          *Engine.ManagedAtlas
	tileset          *Engine.ManagedAtlas
	GameSceneGeneral *PirateScene
	bg               *Engine.GameObject
	floor            *Engine.GameObject
	pl               *Engine.GameObject
	lader            *Engine.GameObject
	splinter         *Engine.GameObject
)

const (
	spr_bg       = 1
	spr_floor    = 2
	spr_lader    = 3
	spr_splinter = 4
)

func CheckError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

type PirateScene struct {
	*Engine.SceneData
	Layer1     *Engine.GameObject
	Layer2     *Engine.GameObject
	Layer3     *Engine.GameObject
	Layer4     *Engine.GameObject
	Background *Engine.GameObject
}

func init() {
	Engine.Title = "PirateLand"
}
func (s *PirateScene) SceneBase() *Engine.SceneData {
	return s.SceneData
}
func (s *PirateScene) Load() {

	Engine.Space.Gravity.Y = -100
	Engine.Space.Iterations = 10
	s.LoadTextures()

	Layer1 := Engine.NewGameObject("Layer1")
	Layer2 := Engine.NewGameObject("Layer2")
	Layer3 := Engine.NewGameObject("Layer3")
	Layer4 := Engine.NewGameObject("Layer4")
	Background := Engine.NewGameObject("Background")

	s.Layer1 = Layer1
	s.Layer2 = Layer2
	s.Layer3 = Layer3
	s.Layer4 = Layer4
	s.Background = Background

	rand.Seed(time.Now().UnixNano())

	GameSceneGeneral = s

	s.Camera = Engine.NewCamera()
	cam := Engine.NewGameObject("Camera")
	cam.AddComponent(s.Camera)

	cam.Transform().SetScalef(1, 1)

	bg = Engine.NewGameObject("bg")
	bg.AddComponent(Engine.NewSprite2(atlas.Texture, Engine.IndexUV(atlas, spr_bg)))
	bg.Transform().SetScalef(2000, 1800)
	bg.Transform().SetPositionf(0, 0)
	bg.Transform().SetParent2(s.Background)

	splinter = Engine.NewGameObject("Splinter")
	splinter.AddComponent(Engine.NewSprite2(atlas.Texture, Engine.IndexUV(atlas, spr_splinter)))
	splinter.Transform().SetWorldScalef(100, 30)
	splinter.AddComponent(Engine.NewPhysics(true, 1, 1))
	splinter.Physics.Shape.IsSensor = true
	splinter.Tag = "splinter"
	for i := 0; i < 1; i++ {
		slc := splinter.Clone()
		slc.Transform().SetParent2(s.Layer3)
		slc.Transform().SetWorldPositionf(230, 130)
	}
	uvs, ind := Engine.AnimatedGroupUVs(plAtlas, "player_walk", "player_stand", "player_attack", "player_jump", "player_bend", "player_hit")

	pl = Engine.NewGameObject("Player")
	pl.AddComponent(Engine.NewSprite3(plAtlas.Texture, uvs))
	pl.Sprite.BindAnimations(ind)
	pl.Sprite.AnimationSpeed = 10
	pl.Transform().SetWorldPositionf(100, 180)
	pl.Transform().SetWorldScalef(50, 100)
	pl.Transform().SetParent2(Layer2)
	pl.AddComponent(NewPlayer())
	pl.AddComponent(Components.NewSmoothFollow(nil, 1, 30))
	pl.AddComponent(Engine.NewPhysics(false, 1, 1))

	lader = Engine.NewGameObject("lader")
	lader.AddComponent(Engine.NewSprite2(atlas.Texture, Engine.IndexUV(atlas, spr_lader)))
	lader.Transform().SetWorldScalef(60, 100)

	lader.Tag = "lader"

	for i := 0; i < 1; i++ {
		lc := lader.Clone()
		lc.Transform().SetParent2(s.Layer3)
		lc.Transform().SetWorldPositionf(100, 100)
	}

	uvs, ind = Engine.AnimatedGroupUVs(tileset, "ground")
	floor = Engine.NewGameObject("floor")
	floor.AddComponent(Engine.NewSprite3(tileset.Texture, uvs))
	floor.Sprite.BindAnimations(ind)
	floor.Sprite.AnimationSpeed = 0
	floor.Transform().SetWorldScalef(100, 100)
	floor.AddComponent(Engine.NewPhysics(true, 1, 1))

	for i := 0; i < 10; i++ {
		f := floor.Clone()
		var h float32 = 50.0

		f.Transform().SetParent2(s.Layer3)
		d := 4
		m := i % 5
		if m == 0 {
			d = 3
		} else if m == 4 {
			d = 5
		}
		if i >= 5 {
			d += 10
			h -= 100
		}
		f.Transform().SetWorldPositionf(float32(i%5)*100, h)
		f.Sprite.SetAnimationIndex(d)
	}
	s.AddGameObject(cam)
	s.AddGameObject(s.Layer1)
	s.AddGameObject(s.Layer2)
	s.AddGameObject(s.Layer3)
	s.AddGameObject(s.Layer4)
	s.AddGameObject(s.Background)

}
func (s *PirateScene) LoadTextures() {
	atlas = Engine.NewManagedAtlas(2048, 1024)
	plAtlas = Engine.NewManagedAtlas(2048, 1024)
	tileset = Engine.NewManagedAtlas(2048, 1024)
	CheckError(atlas.LoadImage("./data/background/backGame.png", spr_bg))
	CheckError(atlas.LoadImage("./data/objects/lader.png", spr_lader))
	CheckError(atlas.LoadImage("./data/objects/splinter.png", spr_splinter))
	CheckError(plAtlas.LoadGroupSheet("./data/player/player_walk.png", 187, 338, 4))
	CheckError(plAtlas.LoadGroupSheet("./data/player/player_stand.png", 187, 338, 1))
	CheckError(plAtlas.LoadGroupSheet("./data/player/player_attack.png", 249, 340, 9))
	CheckError(plAtlas.LoadGroupSheet("./data/player/player_jump.png", 236, 338, 1))
	CheckError(plAtlas.LoadGroupSheet("./data/player/player_bend.png", 188, 259, 1))
	CheckError(plAtlas.LoadGroupSheet("./data/player/player_hit.png", 206, 334, 1))
	CheckError(tileset.LoadGroupSheet("./data/tileset/ground.png", 32, 32, 110))
	atlas.BuildAtlas()
	atlas.BuildMipmaps()
	atlas.SetFiltering(Engine.MipMapLinearNearest, Engine.Nearest)
	atlas.Texture.SetReadOnly()

	plAtlas.BuildAtlas()
	plAtlas.BuildMipmaps()
	plAtlas.SetFiltering(Engine.MipMapLinearNearest, Engine.Nearest)
	plAtlas.Texture.SetReadOnly()

	tileset.BuildAtlas()
	tileset.BuildMipmaps()
	tileset.SetFiltering(Engine.MipMapLinearNearest, Engine.Nearest)
	tileset.Texture.SetReadOnly()
}
func (s *PirateScene) New() Engine.Scene {
	gs := new(PirateScene)
	gs.SceneData = Engine.NewScene("PirateScene")
	return gs
}
