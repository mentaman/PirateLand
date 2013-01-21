package Game

import (
	"fmt"
	"github.com/vova616/GarageEngine/Engine"
	"math/rand"
	"time"
)

var (
	atlas            *Engine.ManagedAtlas
	plAtlas          *Engine.ManagedAtlas
	GameSceneGeneral *PirateScene
	bg               *Engine.GameObject
	floor            *Engine.GameObject
	pl               *Engine.GameObject
)

const (
	spr_bg    = 1
	spr_floor = 2
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
	Background *Engine.GameObject
}

func init() {
	Engine.Title = "PirateLand"
}
func (s *PirateScene) SceneBase() *Engine.SceneData {
	return s.SceneData
}
func (s *PirateScene) Load() {
	s.LoadTextures()

	Layer1 := Engine.NewGameObject("Layer1")
	Background := Engine.NewGameObject("Background")

	s.Layer1 = Layer1
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

	uvs, ind := Engine.AnimatedGroupUVs(plAtlas, "player_walk", "player_stand", "player_attack", "player_jump")
	pl = Engine.NewGameObject("Player")
	pl.AddComponent(Engine.NewSprite3(plAtlas.Texture, uvs))
	pl.Sprite.BindAnimations(ind)
	pl.Sprite.AnimationSpeed = 10
	pl.Transform().SetPositionf(100, 180)
	pl.Transform().SetScalef(50, 100)
	pl.Transform().SetParent2(cam)
	pl.AddComponent(NewPlayer())
	pl.AddComponent(Engine.NewPhysics(false, 1, 1))
	floor = Engine.NewGameObject("floor")
	floor.AddComponent(Engine.NewSprite2(atlas.Texture, Engine.IndexUV(atlas, spr_floor)))
	floor.Transform().SetScalef(100, 100)
	floor.AddComponent(Engine.NewPhysics(true, 1, 1))

	for i := 0; i < 5; i++ {
		f := floor.Clone()
		f.Transform().SetPositionf(float32(i)*350, 50)
		f.Transform().SetParent2(s.Layer1)
	}
	s.AddGameObject(cam)
	s.AddGameObject(s.Layer1)
	s.AddGameObject(s.Background)

}
func (s *PirateScene) LoadTextures() {
	atlas = Engine.NewManagedAtlas(2048, 1024)
	plAtlas = Engine.NewManagedAtlas(2048, 1024)
	CheckError(atlas.LoadImage("./data/backgame.png", spr_bg))
	CheckError(plAtlas.LoadGroupSheet("./data/player_walk.png", 187, 338, 4))
	CheckError(plAtlas.LoadGroupSheet("./data/player_stand.png", 187, 338, 1))
	CheckError(plAtlas.LoadGroupSheet("./data/player_attack.png", 249, 340, 9))
	CheckError(plAtlas.LoadGroupSheet("./data/player_jump.png", 187, 338, 1))
	CheckError(atlas.LoadImage("./data/wall1.png", spr_floor))
	atlas.BuildAtlas()
	atlas.BuildMipmaps()
	atlas.SetFiltering(Engine.MipMapLinearNearest, Engine.Nearest)
	atlas.Texture.SetReadOnly()

	plAtlas.BuildAtlas()
	plAtlas.BuildMipmaps()
	plAtlas.SetFiltering(Engine.MipMapLinearNearest, Engine.Nearest)
	plAtlas.Texture.SetReadOnly()
}
func (s *PirateScene) New() Engine.Scene {
	gs := new(PirateScene)
	gs.SceneData = Engine.NewScene("PirateScene")
	return gs
}
