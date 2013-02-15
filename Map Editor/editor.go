package main

import (
	"fmt"
	"github.com/mentaman/PirateLand/Game/Fonts"

	"github.com/mentaman/PirateLand/Game/Background"
	"github.com/vova616/garageEngine/engine"
)

var (
	atlas       *engine.ManagedAtlas
	objControll *ObjController
	Scene       *MapEditor
	obj         *engine.GameObject
	cam         *engine.GameObject
	Layer1      *engine.GameObject
	sprites     []engine.ID = []engine.ID{"player_walk", "chest", "enemy_walk", "ground"}
	background  *engine.GameObject
)

func CheckError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

type MapEditor struct {
	*engine.SceneData
}

func init() {
}
func (s *MapEditor) SceneBase() *engine.SceneData {
	return s.SceneData
}
func (s *MapEditor) Load() {
	Scene = s
	LoadTextures()
	engine.SetTitle("Map Editor")
	Fonts.ArialFont2, _ = engine.NewFont("../data/Fonts/arial.ttf", 24)
	Fonts.ArialFont2.Texture.SetReadOnly()

	Layer1 = engine.NewGameObject("Layer1")
	background = engine.NewGameObject("Background")

	s.Camera = engine.NewCamera()
	cam = engine.NewGameObject("Camera")
	cam.AddComponent(s.Camera)
	cam.Transform().SetScalef(1, 1)

	mouse := engine.NewGameObject("Mouse")
	mouse.AddComponent(engine.NewMouse())
	mouse.AddComponent(NewMouseController())
	mouse.Transform().SetParent2(cam)

	uvs, ind := engine.AnimatedGroupUVs(atlas, sprites...)
	obj = engine.NewGameObject("Object")
	obj.AddComponent(engine.NewSprite3(atlas.Texture, uvs))

	obj.AddComponent(engine.NewPhysics(false, 1, 1))
	obj.Sprite.BindAnimations(ind)
	obj.Sprite.AnimationSpeed = 0
	obj.AddComponent(NewObject())

	objC := engine.NewGameObject("objController")
	objControll = objC.AddComponent(NewObjController()).(*ObjController)
	objC.Transform().SetParent2(cam)

	Background.Create()
	Background.Object.Transform().SetParent2(background)

	s.AddGameObject(cam)
	s.AddGameObject(Layer1)
	s.AddGameObject(background)

}
func LoadTextures() {
	atlas = engine.NewManagedAtlas(2048, 1024)

	Background.Atlas = engine.NewManagedAtlas(2048, 1024)
	CheckError(Background.Atlas.LoadImageID("../data/background/backGame.png", Background.Spr_bg))
	// 
	// CheckError(Objects.ObjectsAtlas.LoadImageID("./data/objects/lader.png", Objects.Spr_lader))
	// CheckError(Objects.ObjectsAtlas.LoadImageID("./data/objects/splinter.png", Objects.Spr_splinter))
	// CheckError(Objects.ObjectsAtlas.LoadImageID("./data/objects/box.png", Objects.Spr_box))
	// CheckError(Player.ChudAtlas.LoadImageID("./data/bar/chud.png", Player.Spr_chud))
	// CheckError(Player.ChudAtlas.LoadImageID("./data/bar/hpBar.png", Player.Spr_chudHp))
	// CheckError(Player.ChudAtlas.LoadImageID("./data/bar/cpBar.png", Player.Spr_chudCp))
	// CheckError(Player.ChudAtlas.LoadImageID("./data/bar/expBar.png", Player.Spr_chudExp))

	// CheckError(Objects.Atlas.LoadImageID("./data/items/Coin.png", Objects.Spr_coin))
	// CheckError(Objects.Atlas.LoadImageID("./data/items/Coin10.png", Objects.Spr_coin10))
	// CheckError(Objects.Atlas.LoadImageID("./data/items/Daimond.png", Objects.Spr_diamond))
	// CheckError(Objects.Atlas.LoadImageID("./data/items/spot.png", Objects.Spr_spot))
	// CheckError(Objects.Atlas.LoadImageID("./data/items/bigSpot.png", Objects.Spr_bigspot))
	// CheckError(Player.Atlas.LoadImageID("./data/Level/scroll.png", Player.Spr_scroll))

	e, id := atlas.LoadGroupSheet("../data/items/chest.png", 41, 54, 1)
	CheckError(e)

	e, id = atlas.LoadGroupSheet("../data/player/player_walk.png", 187, 338, 1)
	CheckError(e)

	e, id = atlas.LoadGroupSheet("../data/Enemy/enemy_walk.png", 187, 338, 2)
	CheckError(e)

	e, id = atlas.LoadGroupSheet("../data/tileset/ground.png", 32, 32, 110)
	CheckError(e)

	_ = id

	Background.Atlas.BuildAtlas()
	Background.Atlas.BuildMipmaps()
	Background.Atlas.SetFiltering(engine.MipMapLinearNearest, engine.Nearest)
	Background.Atlas.Texture.SetReadOnly()

	atlas.BuildAtlas()
	atlas.BuildMipmaps()
	atlas.SetFiltering(engine.MipMapLinearNearest, engine.Nearest)
	atlas.Texture.SetReadOnly()

}
func (s *MapEditor) New() engine.Scene {
	gs := new(MapEditor)
	gs.SceneData = engine.NewScene("Map Editor")
	return gs
}
