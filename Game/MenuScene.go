package Game

import (
	"github.com/vova616/GarageEngine/Engine"
	"github.com/vova616/GarageEngine/Engine/Components"
)

type MenuScene struct {
	*Engine.SceneData
	layerButtons    *Engine.GameObject
	layerBackground *Engine.GameObject
}

var (
	menuAtlas *Engine.ManagedAtlas
	mbg       *Engine.GameObject
)

const (
	spr_menuback  = 1
	spr_menuexit  = 2
	spr_menunew   = 3
	spr_menuhowTo = 4
	spr_menuload  = 5
)

func (s *MenuScene) SceneBase() *Engine.SceneData {
	return s.SceneData
}
func (s *MenuScene) Load() {
	GameSceneGeneral = s
	LoadTextures()
	s.Camera = Engine.NewCamera()
	cam := Engine.NewGameObject("Camera")
	cam.AddComponent(s.Camera)

	cam.Transform().SetScalef(1, 1)

	mouse := Engine.NewGameObject("Mouse")
	mouse.AddComponent(Engine.NewMouse())
	mouse.Transform().SetParent2(cam)

	layerButtons := Engine.NewGameObject("LayerButton")
	layerBg := Engine.NewGameObject("LayerBackground")

	s.layerBackground = layerBg
	s.layerButtons = layerButtons

	mbg = Engine.NewGameObject("mbg")
	mbg.AddComponent(Engine.NewSprite2(menuAtlas.Texture, Engine.IndexUV(menuAtlas, spr_menuback)))
	mbg.Transform().SetWorldScalef(float32(Engine.Width), float32(Engine.Height))
	mbg.Transform().SetWorldPositionf(float32(Engine.Width)/2, float32(Engine.Height)/2)
	mbg.Transform().SetParent2(s.layerBackground)

	newGame := Engine.NewGameObject("bng")
	newGame.AddComponent(Engine.NewPhysics(false, 1, 1))
	newGame.Physics.Shape.IsSensor = true
	newGame.Transform().SetWorldScalef(100, 100)
	newGame.Transform().SetWorldPositionf(100, 100)
	newGame.Transform().SetParent2(s.layerButtons)
	newGame.AddComponent(Engine.NewSprite2(menuAtlas.Texture, Engine.IndexUV(menuAtlas, spr_menunew)))
	newGame.AddComponent(Components.NewUIButton(func() {
		Engine.LoadScene(Ps)
	}, func(on bool) {
		if on {
			newGame.Sprite.Color = Engine.Vector{1, 0.3, 0.2}
		} else {
			newGame.Sprite.Color = Engine.Vector{1, 1, 1}
		}
	}))

	s.AddGameObject(cam)
	s.AddGameObject(s.layerButtons)
	s.AddGameObject(s.layerBackground)

}
func (s *MenuScene) New() Engine.Scene {
	gs := new(MenuScene)
	gs.SceneData = Engine.NewScene("MenuSceneScene")
	return gs
}
