package Game

import (
	"fmt"
	"github.com/vova616/GarageEngine/Engine"
	"github.com/vova616/GarageEngine/Engine/Components"
	"math/rand"
	"time"
)

var (
	atlas    *Engine.ManagedAtlas
	plAtlas  *Engine.ManagedAtlas
	tileset  *Engine.ManagedAtlas
	bg       *Engine.GameObject
	floor    *Engine.GameObject
	pl       *Engine.GameObject
	lader    *Engine.GameObject
	splinter *Engine.GameObject
	Ps       *PirateScene
	box      *Engine.GameObject
	ch       *Chud

	up         *Engine.GameObject
	cam        *Engine.GameObject
	Layer1     *Engine.GameObject
	Layer2     *Engine.GameObject
	Layer3     *Engine.GameObject
	Layer4     *Engine.GameObject
	Background *Engine.GameObject

	ArialFont2 *Engine.Font
)

const (
	spr_bg       = 1
	spr_floor    = 2
	spr_lader    = 3
	spr_splinter = 4
	spr_box      = 5
	spr_chud     = 6
	spr_chudHp   = 7
	spr_chudCp   = 8
	spr_chudExp  = 9
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
}

func init() {
	Engine.Title = "PirateLand"
}
func (s *PirateScene) SceneBase() *Engine.SceneData {
	return s.SceneData
}
func (s *PirateScene) Load() {
	Ps = s
	LoadTextures()

	ArialFont2, _ = Engine.NewFont("./data/Fonts/arial.ttf", 24)
	ArialFont2.Texture.SetReadOnly()

	Engine.Space.Gravity.Y = -100
	Engine.Space.Iterations = 10

	Layer1 = Engine.NewGameObject("Layer1")
	Layer2 = Engine.NewGameObject("Layer2")
	Layer3 = Engine.NewGameObject("Layer3")
	Layer4 = Engine.NewGameObject("Layer4")
	Background = Engine.NewGameObject("Background")

	rand.Seed(time.Now().UnixNano())

	s.Camera = Engine.NewCamera()
	cam = Engine.NewGameObject("Camera")
	cam.AddComponent(s.Camera)

	cam.Transform().SetScalef(1, 1)

	up = Engine.NewGameObject("up")
	up.Transform().SetParent2(cam)

	bg = Engine.NewGameObject("bg")
	bg.AddComponent(Engine.NewSprite2(atlas.Texture, Engine.IndexUV(atlas, spr_bg)))
	bg.Transform().SetScalef(2000, 1800)
	bg.Transform().SetPositionf(0, 0)
	bg.Transform().SetParent2(Background)

	splinter = Engine.NewGameObject("Splinter")
	splinter.AddComponent(Engine.NewSprite2(atlas.Texture, Engine.IndexUV(atlas, spr_splinter)))
	splinter.Transform().SetWorldScalef(100, 30)
	splinter.AddComponent(Engine.NewPhysics(true, 1, 1))
	splinter.Physics.Shape.IsSensor = true
	splinter.Tag = "splinter"
	for i := 0; i < 1; i++ {
		slc := splinter.Clone()
		slc.Transform().SetParent2(Layer3)
		slc.Transform().SetWorldPositionf(230, 130)
	}
	uvs, ind := Engine.AnimatedGroupUVs(plAtlas, "player_walk", "player_stand", "player_attack", "player_jump", "player_bend", "player_hit", "player_climb")

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
	pl.Physics.Shape.SetFriction(0.7)
	pl.Physics.Shape.SetElasticity(0.2)
	box = Engine.NewGameObject("box")
	box.AddComponent(Engine.NewSprite2(atlas.Texture, Engine.IndexUV(atlas, spr_box)))
	box.Transform().SetWorldScalef(40, 40)
	box.AddComponent(Engine.NewPhysics(false, 1, 1))
	box.Physics.Shape.SetFriction(1)
	for i := 0; i < 1; i++ {
		bc := box.Clone()
		bc.Transform().SetParent2(Layer3)
		bc.Transform().SetWorldPositionf(30, 150)
	}

	lader = Engine.NewGameObject("lader")
	lader.AddComponent(Engine.NewSprite2(atlas.Texture, Engine.IndexUV(atlas, spr_lader)))
	lader.Transform().SetWorldScalef(60, 100)
	lader.AddComponent(Engine.NewPhysics(true, 1, 1))
	lader.Physics.Shape.IsSensor = true
	lader.Physics.Shape.SetFriction(2)
	lader.Tag = "lader"

	for i := 0; i < 1; i++ {
		lc := lader.Clone()
		lc.Transform().SetParent2(Layer3)
		lc.Transform().SetWorldPositionf(150, 150)
	}

	uvs, ind = Engine.AnimatedGroupUVs(tileset, "ground")
	floor = Engine.NewGameObject("floor")
	floor.AddComponent(Engine.NewSprite3(tileset.Texture, uvs))
	floor.Sprite.BindAnimations(ind)
	floor.Sprite.AnimationSpeed = 0
	floor.Transform().SetWorldScalef(100, 100)
	floor.AddComponent(Engine.NewPhysics(true, 1, 1))

	chud := Engine.NewGameObject("chud")
	chud.AddComponent(Engine.NewSprite2(atlas.Texture, Engine.IndexUV(atlas, spr_chud)))
	ch = chud.AddComponent(NewChud()).(*Chud)
	chud.Transform().SetParent2(cam)
	chud.Transform().SetWorldPositionf(200, 550)
	chud.Transform().SetWorldScalef(100, 100)

	label := Engine.NewGameObject("Label")
	label.Transform().SetParent2(cam)
	label.Transform().SetPositionf(20, float32(Engine.Height)-40)
	label.Transform().SetScalef(20, 20)

	Hp := Engine.NewGameObject("hpBar")
	Hp.GameObject().AddComponent(Engine.NewSprite2(atlas.Texture, Engine.IndexUV(atlas, spr_chudHp)))
	Hp.GameObject().Sprite.SetAlign(Engine.AlignLeft)
	Hp.GameObject().Transform().SetWorldPosition(Engine.Vector{235, 580, 0})
	Hp.GameObject().Transform().SetWorldScalef(17, 20)
	Hp.Transform().SetParent2(up)
	ch.Hp = (Hp.AddComponent(NewBar(17))).(*Bar)

	txt2 := label.AddComponent(NewTestBox(func(tx *TestTextBox) {
		Hp.Transform().SetPositionf(235, float32(tx.V))
	})).(*TestTextBox)
	txt2.SetAlign(Engine.AlignLeft)

	for i := 0; i < 10; i++ {
		f := floor.Clone()
		var h float32 = 50.0

		f.Transform().SetParent2(Layer3)
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
	s.AddGameObject(up)
	s.AddGameObject(cam)
	s.AddGameObject(Layer1)
	s.AddGameObject(Layer2)
	s.AddGameObject(Layer3)
	s.AddGameObject(Layer4)
	s.AddGameObject(Background)

}
func LoadTextures() {
	atlas = Engine.NewManagedAtlas(2048, 1024)
	plAtlas = Engine.NewManagedAtlas(2048, 1024)
	tileset = Engine.NewManagedAtlas(2048, 1024)
	menuAtlas = Engine.NewManagedAtlas(2048, 1024)
	CheckError(atlas.LoadImage("./data/background/backGame.png", spr_bg))
	CheckError(atlas.LoadImage("./data/objects/lader.png", spr_lader))
	CheckError(atlas.LoadImage("./data/objects/splinter.png", spr_splinter))
	CheckError(atlas.LoadImage("./data/objects/box.png", spr_box))
	CheckError(atlas.LoadImage("./data/bar/chud.png", spr_chud))
	CheckError(atlas.LoadImage("./data/bar/hpBar.png", spr_chudHp))
	CheckError(atlas.LoadImage("./data/bar/cpBar.png", spr_chudCp))
	CheckError(atlas.LoadImage("./data/bar/expBar.png", spr_chudExp))
	CheckError(plAtlas.LoadGroupSheet("./data/player/player_walk.png", 187, 338, 4))
	CheckError(plAtlas.LoadGroupSheet("./data/player/player_stand.png", 187, 338, 1))
	CheckError(plAtlas.LoadGroupSheet("./data/player/player_attack.png", 249, 340, 9))
	CheckError(plAtlas.LoadGroupSheet("./data/player/player_jump.png", 236, 338, 1))
	CheckError(plAtlas.LoadGroupSheet("./data/player/player_bend.png", 188, 259, 1))
	CheckError(plAtlas.LoadGroupSheet("./data/player/player_climb.png", 236, 363, 2))
	CheckError(plAtlas.LoadGroupSheet("./data/player/player_hit.png", 206, 334, 1))
	CheckError(tileset.LoadGroupSheet("./data/tileset/ground.png", 32, 32, 110))
	CheckError(menuAtlas.LoadImage("./data/menu/menuback.png", spr_menuback))
	CheckError(menuAtlas.LoadImage("./data/menu/exit.png", spr_menuexit))
	CheckError(menuAtlas.LoadImage("./data/menu/newgame.png", spr_menunew))
	CheckError(menuAtlas.LoadImage("./data/menu/load.png", spr_menuload))
	CheckError(menuAtlas.LoadImage("./data/menu/howTo.png", spr_menuhowTo))
	atlas.BuildAtlas()
	atlas.BuildMipmaps()
	atlas.SetFiltering(Engine.MipMapLinearNearest, Engine.Nearest)
	atlas.Texture.SetReadOnly()

	plAtlas.BuildAtlas()
	plAtlas.BuildMipmaps()
	plAtlas.SetFiltering(Engine.MipMapLinearNearest, Engine.Nearest)
	plAtlas.Texture.SetReadOnly()

	menuAtlas.BuildAtlas()
	menuAtlas.BuildMipmaps()
	menuAtlas.SetFiltering(Engine.MipMapLinearNearest, Engine.Nearest)
	menuAtlas.Texture.SetReadOnly()

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
