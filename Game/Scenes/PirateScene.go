package Game

import (
	"fmt"
	"github.com/mentaman/PirateLand/Game/Fonts"
	"github.com/mentaman/PirateLand/Game/Player"

	"github.com/mentaman/PirateLand/Game/Background"
	"github.com/mentaman/PirateLand/Game/Enemy"
	"github.com/mentaman/PirateLand/Game/GUI"
	"github.com/mentaman/PirateLand/Game/Objects"
	"github.com/vova616/garageEngine/engine"
	"math/rand"
	"time"
)

var (
	atlas *engine.ManagedAtlas
	bg    *engine.GameObject
	en    *engine.GameObject
	spot  *engine.GameObject

	Ps *PirateScene

	up     *engine.GameObject
	cam    *engine.GameObject
	Layer1 *engine.GameObject
	Layer2 *engine.GameObject
	Layer3 *engine.GameObject
	Layer4 *engine.GameObject

	background *engine.GameObject
)

func CheckError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

type PirateScene struct {
	*engine.SceneData
}

func init() {
}
func (s *PirateScene) SceneBase() *engine.SceneData {
	return s.SceneData
}
func (s *PirateScene) Load() {
	Ps = s
	LoadTextures()
	engine.SetTitle("PirateLand")
	Fonts.ArialFont2, _ = engine.NewFont("./data/Fonts/arial.ttf", 24)
	Fonts.ArialFont2.Texture.SetReadOnly()

	engine.Space.Gravity.Y = -100
	engine.Space.Iterations = 10

	Layer1 = engine.NewGameObject("Layer1")
	Layer2 = engine.NewGameObject("Layer2")
	Layer3 = engine.NewGameObject("Layer3")
	Layer4 = engine.NewGameObject("Layer4")
	background = engine.NewGameObject("Background")

	rand.Seed(time.Now().UnixNano())

	s.Camera = engine.NewCamera()
	cam = engine.NewGameObject("Camera")
	cam.AddComponent(s.Camera)

	cam.Transform().SetScalef(1, 1)

	up = engine.NewGameObject("up")
	up.Transform().SetParent2(cam)

	Objects.CreateObjects()
	Background.Create()
	Background.Object.Transform().SetParent2(background)

	for i := 0; i < 1; i++ {
		slc := Objects.Splinter.Clone()
		slc.Transform().SetParent2(Layer3)
		slc.Transform().SetWorldPositionf(230, 130)
	}

	Player.CreateChud()
	Player.Ch.Transform().SetParent2(cam)
	Player.CreatePlayer()
	Player.Pl.Transform().SetParent2(Layer2)

	Enemy.CreateEnemy()
	sd := GUI.NewBar(10)
	for i := 0; i < 2; i++ {
		ec := Enemy.Regular.Clone()
		hpB := Enemy.HpBar.Clone()
		hpBd := hpB.ComponentTypeOfi(sd).(*GUI.Bar)

		ec.Transform().SetWorldPositionf(200+rand.Float32(), 110)
		ec.AddComponent(Enemy.NewEnemy(hpBd))
		ec.Transform().SetParent2(Layer2)

		ec.Sprite.AnimationSpeed = 10
		hpB.Transform().SetParent2(up)
	}

	for i := 0; i < 1; i++ {
		bc := Objects.Box.Clone()
		bc.Transform().SetParent2(Layer3)
		bc.Transform().SetWorldPositionf(30, 150)
	}

	for i := 0; i < 1; i++ {
		lc := Objects.Lader.Clone()
		lc.Transform().SetParent2(Layer3)
		lc.Transform().SetWorldPositionf(150, 150)
	}

	label := engine.NewGameObject("Label")
	label.Transform().SetParent2(cam)
	label.Transform().SetPositionf(20, float32(engine.Height)-40)
	label.Transform().SetScalef(20, 20)

	for i := 0; i < 2; i++ {
		s := Objects.ChestO.Clone()
		s.Transform().SetWorldPositionf(300+20*float32(i)+float32(rand.Int()%300), 150)
		s.Transform().SetParent2(Layer3)
	}

	Player.Ch.Hp.Transform().SetParent2(up)
	Player.Ch.Cp.Transform().SetParent2(up)
	Player.Ch.Exp.Transform().SetParent2(up)
	Player.Ch.Exp.Start()
	Player.Ch.Exp.SetValue(0, 100)
	Player.Ch.Money.Transform().SetParent2(cam)
	Player.Ch.Level.Transform().SetParent2(cam)
	Player.PlComp.MenuScene = func() {
		engine.LoadScene(MenuSceneG)
	}
	txt2 := label.AddComponent(GUI.NewTestBox(func(tx *GUI.TestTextBox) {
		Player.Ch.Cp.Transform().SetPositionf(156, float32(tx.V))
	})).(*GUI.TestTextBox)
	txt2.SetAlign(engine.AlignLeft)
	for i := 0; i < 10; i++ {
		f := Objects.Floor.Clone()
		f.Sprite.SetAnimation("ground")
		f.Transform().SetWorldPositionf(float32(i)*100, 50)
		f.Sprite.SetAnimationIndex(4)
		f.Transform().SetParent2(Layer3)
	}
	s.AddGameObject(up)
	s.AddGameObject(cam)
	s.AddGameObject(Layer1)
	s.AddGameObject(Layer2)
	s.AddGameObject(Layer3)
	s.AddGameObject(Layer4)
	s.AddGameObject(background)

}
func LoadTextures() {
	atlas = engine.NewManagedAtlas(2048, 1024)
	Background.Atlas = engine.NewManagedAtlas(2048, 1024)
	Objects.Atlas = engine.NewManagedAtlas(2048, 1024)
	Objects.ObjectsAtlas = engine.NewManagedAtlas(2048, 1024)
	Player.Atlas = engine.NewManagedAtlas(2048, 1024)
	Enemy.Atlas = engine.NewManagedAtlas(2048, 1024)
	Objects.Tileset = engine.NewManagedAtlas(2048, 1024)
	menuAtlas = engine.NewManagedAtlas(2048, 1024)
	Player.ChudAtlas = engine.NewManagedAtlas(2048, 1024)

	CheckError(Background.Atlas.LoadImageID("./data/background/backGame.png", Background.Spr_bg))
	CheckError(Objects.ObjectsAtlas.LoadImageID("./data/objects/lader.png", Objects.Spr_lader))
	CheckError(Objects.ObjectsAtlas.LoadImageID("./data/objects/splinter.png", Objects.Spr_splinter))
	CheckError(Objects.ObjectsAtlas.LoadImageID("./data/objects/box.png", Objects.Spr_box))
	CheckError(Player.ChudAtlas.LoadImageID("./data/bar/chud.png", Player.Spr_chud))
	CheckError(Player.ChudAtlas.LoadImageID("./data/bar/hpBar.png", Player.Spr_chudHp))
	CheckError(Player.ChudAtlas.LoadImageID("./data/bar/cpBar.png", Player.Spr_chudCp))
	CheckError(Player.ChudAtlas.LoadImageID("./data/bar/expBar.png", Player.Spr_chudExp))
	Objects.ObjectsAtlas.LoadGroupSheet("./data/items/chest.png", 41, 54, 4)
	CheckError(Objects.Atlas.LoadImageID("./data/items/Coin.png", Objects.Spr_coin))
	CheckError(Objects.Atlas.LoadImageID("./data/items/Coin10.png", Objects.Spr_coin10))
	CheckError(Objects.Atlas.LoadImageID("./data/items/Daimond.png", Objects.Spr_diamond))
	CheckError(Objects.Atlas.LoadImageID("./data/items/spot.png", Objects.Spr_spot))
	CheckError(Objects.Atlas.LoadImageID("./data/items/bigSpot.png", Objects.Spr_bigspot))
	CheckError(Player.Atlas.LoadImageID("./data/Level/scroll.png", Player.Spr_scroll))

	e, id := Player.Atlas.LoadGroupSheet("./data/player/player_walk.png", 187, 338, 4)
	CheckError(e)
	e, id = Player.Atlas.LoadGroupSheet("./data/player/player_stand.png", 187, 338, 1)
	CheckError(e)
	e, id = Player.Atlas.LoadGroupSheet("./data/player/player_attack.png", 249, 340, 9)
	CheckError(e)
	e, id = Player.Atlas.LoadGroupSheet("./data/player/player_jump.png", 236, 338, 1)
	CheckError(e)
	e, id = Player.Atlas.LoadGroupSheet("./data/player/player_bend.png", 188, 259, 1)
	CheckError(e)
	e, id = Player.Atlas.LoadGroupSheet("./data/player/player_climb.png", 236, 363, 2)
	CheckError(e)
	e, id = Player.Atlas.LoadGroupSheet("./data/player/player_hit.png", 206, 334, 1)
	CheckError(e)

	e, id = Enemy.Atlas.LoadGroupSheet("./data/Enemy/enemy_walk.png", 187, 338, 4)
	CheckError(e)
	e, id = Enemy.Atlas.LoadGroupSheet("./data/Enemy/enemy_stand.png", 187, 338, 1)
	CheckError(e)
	e, id = Enemy.Atlas.LoadGroupSheet("./data/Enemy/enemy_attack.png", 249, 340, 9)
	CheckError(e)
	e, id = Enemy.Atlas.LoadGroupSheet("./data/Enemy/enemy_jump.png", 236, 338, 1)
	CheckError(e)
	e, id = Enemy.Atlas.LoadGroupSheet("./data/Enemy/enemy_hit.png", 206, 334, 1)
	CheckError(e)

	e, id = Objects.Tileset.LoadGroupSheet("./data/tileset/ground.png", 32, 32, 110)
	CheckError(e)

	_ = id

	CheckError(menuAtlas.LoadImageID("./data/menu/menuback.png", spr_menuback))
	CheckError(menuAtlas.LoadImageID("./data/menu/exit.png", spr_menuexit))
	CheckError(menuAtlas.LoadImageID("./data/menu/newgame.png", spr_menunew))
	CheckError(menuAtlas.LoadImageID("./data/menu/load.png", spr_menuload))
	CheckError(menuAtlas.LoadImageID("./data/menu/howTo.png", spr_menuhowTo))
	atlas.BuildAtlas()
	atlas.BuildMipmaps()
	atlas.SetFiltering(engine.MipMapLinearNearest, engine.Nearest)
	atlas.Texture.SetReadOnly()

	Background.Atlas.BuildAtlas()
	Background.Atlas.BuildMipmaps()
	Background.Atlas.SetFiltering(engine.MipMapLinearNearest, engine.Nearest)
	Background.Atlas.Texture.SetReadOnly()

	Objects.Atlas.BuildAtlas()
	Objects.Atlas.BuildMipmaps()
	Objects.Atlas.SetFiltering(engine.MipMapLinearNearest, engine.Nearest)
	Objects.Atlas.Texture.SetReadOnly()

	Objects.ObjectsAtlas.BuildAtlas()
	Objects.ObjectsAtlas.BuildMipmaps()
	Objects.ObjectsAtlas.SetFiltering(engine.MipMapLinearNearest, engine.Nearest)
	Objects.ObjectsAtlas.Texture.SetReadOnly()

	Enemy.Atlas.BuildAtlas()
	Enemy.Atlas.BuildMipmaps()
	Enemy.Atlas.SetFiltering(engine.MipMapLinearNearest, engine.Nearest)
	Enemy.Atlas.Texture.SetReadOnly()

	Player.ChudAtlas.BuildAtlas()
	Player.ChudAtlas.BuildMipmaps()
	Player.ChudAtlas.SetFiltering(engine.MipMapLinearNearest, engine.Nearest)
	Player.ChudAtlas.Texture.SetReadOnly()

	Player.Atlas.BuildAtlas()
	Player.Atlas.BuildMipmaps()
	Player.Atlas.SetFiltering(engine.MipMapLinearNearest, engine.Nearest)
	Player.Atlas.Texture.SetReadOnly()

	menuAtlas.BuildAtlas()
	menuAtlas.BuildMipmaps()
	menuAtlas.SetFiltering(engine.MipMapLinearNearest, engine.Nearest)
	menuAtlas.Texture.SetReadOnly()

	Objects.Tileset.BuildAtlas()
	Objects.Tileset.BuildMipmaps()
	Objects.Tileset.SetFiltering(engine.MipMapLinearNearest, engine.Nearest)
	Objects.Tileset.Texture.SetReadOnly()
}
func (s *PirateScene) New() engine.Scene {
	gs := new(PirateScene)
	gs.SceneData = engine.NewScene("PirateScene")
	return gs
}
