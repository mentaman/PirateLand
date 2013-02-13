package Game

import (
	"github.com/vova616/garageEngine/engine"
	"github.com/vova616/garageEngine/engine/components"
)

type MenuScene struct {
	*engine.SceneData
	layerButtons    *engine.GameObject
	layerBackground *engine.GameObject
}

var (
	menuAtlas  *engine.ManagedAtlas
	mbg        *engine.GameObject
	MenuSceneG *MenuScene
)

const (
	spr_menuback  = 1
	spr_menuexit  = 2
	spr_menunew   = 3
	spr_menuhowTo = 4
	spr_menuload  = 5
)

func (s *MenuScene) SceneBase() *engine.SceneData {
	return s.SceneData
}
func (s *MenuScene) Load() {
	MenuSceneG = s
	LoadTextures()
	engine.SetTitle("PirateLand - menu")
	s.Camera = engine.NewCamera()
	cam := engine.NewGameObject("Camera")
	cam.AddComponent(s.Camera)

	cam.Transform().SetScalef(1, 1)

	mouse := engine.NewGameObject("Mouse")
	mouse.AddComponent(engine.NewMouse())
	mouse.Transform().SetParent2(cam)

	layerButtons := engine.NewGameObject("LayerButton")
	layerBg := engine.NewGameObject("LayerBackground")

	s.layerBackground = layerBg
	s.layerButtons = layerButtons

	mbg = engine.NewGameObject("mbg")
	mbg.AddComponent(engine.NewSprite2(menuAtlas.Texture, engine.IndexUV(menuAtlas, spr_menuback)))
	mbg.Transform().SetWorldScalef(float32(engine.Width), float32(engine.Height))
	mbg.Transform().SetWorldPositionf(float32(engine.Width)/2, float32(engine.Height)/2)
	mbg.Transform().SetParent2(s.layerBackground)

	newGame := engine.NewGameObject("bng")
	newGame.AddComponent(engine.NewPhysics(false, 1, 1))
	newGame.Physics.Shape.IsSensor = true
	newGame.Transform().SetWorldScalef(100, 100)
	newGame.Transform().SetWorldPositionf(100, 100)
	newGame.Transform().SetParent2(s.layerButtons)
	newGame.AddComponent(engine.NewSprite2(menuAtlas.Texture, engine.IndexUV(menuAtlas, spr_menunew)))
	newGame.AddComponent(components.NewUIButton(func() {
		engine.LoadScene(Ps)
	}, func(on bool) {
		if on {
			newGame.Sprite.Color = engine.Color{1, 0.3, 0.2, 1}
		} else {
			newGame.Sprite.Color = engine.Color{1, 1, 1, 1}
		}
	}))

	s.AddGameObject(cam)
	s.AddGameObject(s.layerButtons)
	s.AddGameObject(s.layerBackground)

}
func (s *MenuScene) New() engine.Scene {
	gs := new(MenuScene)
	gs.SceneData = engine.NewScene("MenuSceneScene")
	return gs
}
