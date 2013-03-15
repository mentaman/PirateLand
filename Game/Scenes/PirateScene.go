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

	"github.com/vova616/garageEngine/engine/components"
	"math/rand"
	"time"
)

var (
	menuAtlas *engine.ManagedAtlas
	mbg       *engine.GameObject
	master    *engine.GameObject
)

const (
	spr_menuback     = 1
	spr_menuexit     = 2
	spr_menunew      = 3
	spr_menuhowTo    = 4
	spr_menuload     = 5
	spr_menucontinue = 6
)

var (
	atlas    *engine.ManagedAtlas
	bg       *engine.GameObject
	en       *engine.GameObject
	spot     *engine.GameObject
	CamLayer *engine.GameObject
	Ps       *PirateScene

	Layer1 *engine.GameObject
	Layer2 *engine.GameObject
	Layer3 *engine.GameObject
	Layer4 *engine.GameObject

	Container  *engine.GameObject
	background *engine.GameObject
	con        bool = false

	gameLoaded = false
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
	PirateCam       *engine.Camera
	MenuCam         *engine.Camera
	layerButtons    *engine.GameObject
	layerBackground *engine.GameObject
}

func init() {
}
func (s *PirateScene) SceneBase() *engine.SceneData {
	return s.SceneData
}
func (s *PirateScene) Load() {
	// if !con {
	// 	s.FakeLoad()
	// } else {
	// 	if Container != nil {
	// 		s.AddGameObject(Container)
	// 	}
	// }
	Ps = s

	LoadTextures()

	engine.SetTitle("PirateLand")
	s.MenuLoad()
}
func (s *PirateScene) MenuLoad() {
	LoadTextures()

	engine.Space.Gravity.Y = 0
	engine.SetTitle("PirateLand - menu")

	s.Camera = engine.NewCamera()
	cam := engine.NewGameObject("Camera")
	cam.AddComponent(s.Camera)
	master = engine.NewGameObject("master")
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
	fmt.Println(engine.Width/2, engine.Height/2)
	mbg.Transform().SetWorldPositionf(0, 0)
	mbg.Transform().SetParent2(s.layerBackground)
	mbg.Transform().SetDepth(-1)
	newGame := engine.NewGameObject("bng")
	newGame.Transform().SetWorldScalef(100, 100)
	newGame.Transform().SetWorldPositionf(-240, -160)
	newGame.Transform().SetParent2(s.layerButtons)

	newGame.AddComponent(engine.NewPhysics(false))
	newGame.Physics.Shape.IsSensor = true
	newGame.AddComponent(engine.NewSprite2(menuAtlas.Texture, engine.IndexUV(menuAtlas, spr_menunew)))
	newGame.AddComponent(components.NewUIButton(func() {
		s.RemoveGameObject(master)
		s.GameLoad()
		con = false
	}, func(on bool) {
		if on {
			newGame.Sprite.Color = engine.Color{1, 0.3, 0.2, 1}
		} else {
			newGame.Sprite.Color = engine.Color{1, 1, 1, 1}
		}
	}))

	continueGame := engine.NewGameObject("bng")
	continueGame.Transform().SetWorldScalef(100, 100)
	continueGame.Transform().SetWorldPositionf(-40, -60)
	continueGame.Transform().SetParent2(s.layerButtons)

	continueGame.AddComponent(engine.NewPhysics(false))
	continueGame.Physics.Shape.IsSensor = true
	continueGame.AddComponent(engine.NewSprite2(menuAtlas.Texture, engine.IndexUV(menuAtlas, spr_menucontinue)))
	if !gameLoaded {
		continueGame.Sprite.Color = engine.Color{0.2, 0.2, 0.2, 1}
	}
	continueGame.AddComponent(components.NewUIButton(func() {

		if gameLoaded {
			s.RemoveGameObject(master)
			s.GameContinue()
			con = false
		}

	}, func(on bool) {
		if gameLoaded {
			if on {
				continueGame.Sprite.Color = engine.Color{1, 0.3, 0.2, 1}
			} else {
				continueGame.Sprite.Color = engine.Color{1, 1, 1, 1}
			}
		} else {
			if on {
				continueGame.Sprite.Color = engine.Color{0.3, 0.3, 0.3, 1}
			} else {
				continueGame.Sprite.Color = engine.Color{0.2, 0.2, 0.2, 1}
			}
		}
	}))

	loadGame := engine.NewGameObject("bng")
	loadGame.Transform().SetWorldScalef(100, 100)
	loadGame.Transform().SetWorldPositionf(160, -160)
	loadGame.Transform().SetParent2(s.layerButtons)

	loadGame.AddComponent(engine.NewPhysics(false))
	loadGame.Physics.Shape.IsSensor = true
	loadGame.AddComponent(engine.NewSprite2(menuAtlas.Texture, engine.IndexUV(menuAtlas, spr_menuload)))
	loadGame.AddComponent(components.NewUIButton(func() {

		if gameLoaded {
			// s.RemoveGameObject(master)
			// s.GameContinue()
			// con = false
		}

	}, func(on bool) {
		if on {
			loadGame.Sprite.Color = engine.Color{1, 0.3, 0.2, 1}
		} else {
			loadGame.Sprite.Color = engine.Color{1, 1, 1, 1}
		}

	}))

	cam.Transform().SetParent2(master)

	s.layerButtons.Transform().SetParent2(master)
	s.layerBackground.Transform().SetParent2(master)
	s.AddGameObject(master)
}
func (s *PirateScene) GameContinue() {
	s.AddGameObject(Container)

	engine.SetTitle("PirateLand")
	engine.Space.Gravity.Y = -100
	engine.Space.Iterations = 10

	s.Camera = s.PirateCam
}

func (s *PirateScene) GameLoad() {
	if Container != nil {
		Container.Destroy()
	}
	gameLoaded = true

	Fonts.ArialFont2, _ = engine.NewFont("./data/Fonts/arial.ttf", 24)
	Fonts.ArialFont2.Texture.SetReadOnly()

	engine.SetTitle("PirateLand")
	engine.Space.Gravity.Y = -100
	engine.Space.Iterations = 10

	Layer1 = engine.NewGameObject("Layer1")

	Container = engine.NewGameObject("Container")
	Layer2 = engine.NewGameObject("Layer2")
	Layer3 = engine.NewGameObject("Layer3")
	Layer4 = engine.NewGameObject("Layer4")
	background = engine.NewGameObject("Background")

	rand.Seed(time.Now().UnixNano())
	s.PirateCam = engine.NewCamera()
	s.Camera = s.PirateCam
	CamLayer = engine.NewGameObject("Camera")
	CamLayer.AddComponent(s.Camera)

	CamLayer.Transform().SetScalef(1, 1)

	Objects.CreateObjects()

	Objects.BirdControll = engine.StartCoroutine(func() {
		for {
			engine.CoSleep(float32(rand.Intn(6)))
			b := Objects.BirdO.Clone()
			b.AddComponent(Objects.NewBird(rand.Int()%2 == 1, float32(rand.Intn(10)+1)))
			b.Transform().SetParent2(Layer2)
			b.Transform().SetDepth(-1)
		}
	})

	Background.Create()
	Background.Object.Transform().SetParent2(background)

	// for i := 0; i < 1; i++ {
	// 	slc := Objects.Splinter.Clone()
	// 	slc.Transform().SetParent2(Layer3)
	// 	slc.Transform().SetWorldPositionf(230, 130)
	// }

	Player.CreateChud()
	Player.Ch.Transform().SetParent2(CamLayer)
	Player.CreatePlayer()
	Player.Pl.Transform().SetParent2(Layer2)

	Player.Sk.Transform().SetParent2(CamLayer)
	Enemy.CreateEnemy()
	// sd := GUI.NewBar(10)
	// for i := 0; i < 2; i++ {
	// 	ec := Enemy.Regular.Clone()
	// 	hpB := Enemy.HpBar.Clone()
	// 	hpBd := hpB.ComponentTypeOfi(sd).(*GUI.Bar)

	// 	ec.Transform().SetWorldPositionf(200+rand.Float32(), 110)
	// 	ec.AddComponent(Enemy.NewEnemy(hpBd))
	// 	ec.Transform().SetParent2(Layer2)

	// 	ec.Sprite.AnimationSpeed = 10
	// 	hpB.Transform().SetParent2(CamLayer)
	// }

	// for i := 0; i < 1; i++ {
	// 	bc := Objects.Box.Clone()
	// 	bc.Transform().SetParent2(Layer3)
	// 	bc.Transform().SetWorldPositionf(30, 150)
	// }

	// for i := 0; i < 1; i++ {
	// 	lc := Objects.Lader.Clone()
	// 	lc.Transform().SetParent2(Layer3)
	// 	lc.Transform().SetWorldPositionf(150, 150)
	// }
	mouse := engine.NewGameObject("Mouse")
	mouse.AddComponent(engine.NewMouse())
	mouse.Transform().SetParent2(CamLayer)

	label := engine.NewGameObject("Label")
	label.Transform().SetParent2(CamLayer)
	label.Transform().SetPositionf(20, float32(engine.Height)-40)
	label.Transform().SetScalef(20, 20)

	// for i := 0; i < 2; i++ {
	// 	s := Objects.ChestO.Clone()
	// 	s.Transform().SetWorldPositionf(300+20*float32(i)+float32(rand.Int()%300), 150)
	// 	s.Transform().SetParent2(Layer3)
	// }

	Player.Ch.Hp.Transform().SetParent2(CamLayer)
	Player.Ch.Cp.Transform().SetParent2(CamLayer)
	Player.Ch.Exp.Transform().SetParent2(CamLayer)
	Player.Ch.Exp.Start()
	Player.Ch.Exp.SetValue(0, 100)
	Player.Ch.Money.Transform().SetParent2(CamLayer)
	Player.Ch.Level.Transform().SetParent2(CamLayer)
	Player.PlComp.MenuScene = func() {
		s.RemoveGameObject(Container)
		s.MenuLoad()
	}
	txt2 := label.AddComponent(GUI.NewTestBox(func(tx *GUI.TestTextBox) {
		Player.Ch.Level.Transform().SetPositionf(float32(tx.V), 530)
	})).(*GUI.TestTextBox)
	txt2.SetAlign(engine.AlignLeft)
	// for i := 0; i < 10; i++ {
	// 	f := Objects.Floor.Clone()
	// 	f.Sprite.SetAnimation("ground")
	// 	f.Transform().SetWorldPositionf(float32(i)*100, 50)
	// 	f.Sprite.SetAnimationIndex(4)
	// 	f.Transform().SetParent2(Layer3)
	// }
	LoadMap("./maps/map.txt")
	CamLayer.Transform().SetParent2(Container)
	Layer1.Transform().SetParent2(Container)
	Layer2.Transform().SetParent2(Container)
	Layer3.Transform().SetParent2(Container)
	Layer4.Transform().SetParent2(Container)
	background.Transform().SetParent2(Container)

	s.AddGameObject(Container)

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
	Player.SkillsAtlas = engine.NewManagedAtlas(2048, 1024)
	CheckError(Background.Atlas.LoadImageID("./data/background/backGame.png", Background.Spr_bg))
	CheckError(Objects.ObjectsAtlas.LoadImageID("./data/objects/lader.png", Objects.Spr_lader))
	CheckError(Objects.ObjectsAtlas.LoadImageID("./data/objects/splinter.png", Objects.Spr_splinter))
	CheckError(Objects.ObjectsAtlas.LoadImageID("./data/objects/box.png", Objects.Spr_box))
	CheckError(Objects.ObjectsAtlas.LoadImageID("./data/objects/bird.png", Objects.Spr_bird))
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
	CheckError(Player.SkillsAtlas.LoadImageID("./data/skills/skillMenu.png", Player.Spr_skillGUI))
	CheckError(Player.SkillsAtlas.LoadImageID("./data/skills/button1.png", Player.Spr_skillHp))
	CheckError(Player.SkillsAtlas.LoadImageID("./data/skills/button2.png", Player.Spr_skillStrong))
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
	CheckError(menuAtlas.LoadImageID("./data/menu/continue.png", spr_menucontinue))
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

	Player.SkillsAtlas.BuildAtlas()
	Player.SkillsAtlas.BuildMipmaps()
	Player.SkillsAtlas.SetFiltering(engine.MipMapLinearNearest, engine.Nearest)
	Player.SkillsAtlas.Texture.SetReadOnly()

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
