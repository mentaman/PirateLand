package Game

import (
	"fmt"
	"github.com/vova616/garageEngine/engine"
	"github.com/vova616/garageEngine/engine/components"
	"math/rand"
	"time"
)

var (
	atlas    *engine.ManagedAtlas
	plAtlas  *engine.ManagedAtlas
	enAtlas  *engine.ManagedAtlas
	tileset  *engine.ManagedAtlas
	bg       *engine.GameObject
	floor    *engine.GameObject
	pl       *engine.GameObject
	en       *engine.GameObject
	lader    *engine.GameObject
	splinter *engine.GameObject
	spot     *engine.GameObject
	Ps       *PirateScene
	box      *engine.GameObject
	chest    *engine.GameObject
	ch       *Chud

	up         *engine.GameObject
	cam        *engine.GameObject
	Layer1     *engine.GameObject
	Layer2     *engine.GameObject
	Layer3     *engine.GameObject
	Layer4     *engine.GameObject
	Background *engine.GameObject

	ArialFont2 *engine.Font
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
	spr_coin     = 10
	spr_coin10   = 11
	spr_diamond  = 12
	spr_spot     = 13
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
	engine.Title = "PirateLand"
}
func (s *PirateScene) SceneBase() *engine.SceneData {
	return s.SceneData
}
func (s *PirateScene) Load() {
	Ps = s
	LoadTextures()

	ArialFont2, _ = engine.NewFont("./data/Fonts/arial.ttf", 24)
	ArialFont2.Texture.SetReadOnly()

	engine.Space.Gravity.Y = -100
	engine.Space.Iterations = 10

	Layer1 = engine.NewGameObject("Layer1")
	Layer2 = engine.NewGameObject("Layer2")
	Layer3 = engine.NewGameObject("Layer3")
	Layer4 = engine.NewGameObject("Layer4")
	Background = engine.NewGameObject("Background")

	rand.Seed(time.Now().UnixNano())

	s.Camera = engine.NewCamera()
	cam = engine.NewGameObject("Camera")
	cam.AddComponent(s.Camera)

	cam.Transform().SetScalef(1, 1)

	up = engine.NewGameObject("up")
	up.Transform().SetParent2(cam)

	bg = engine.NewGameObject("bg")
	bg.AddComponent(engine.NewSprite2(atlas.Texture, engine.IndexUV(atlas, spr_bg)))
	bg.Transform().SetScalef(2000, 1800)
	bg.Transform().SetPositionf(0, 0)
	bg.Transform().SetParent2(Background)

	splinter = engine.NewGameObject("Splinter")
	splinter.AddComponent(engine.NewSprite2(atlas.Texture, engine.IndexUV(atlas, spr_splinter)))
	splinter.Transform().SetWorldScalef(100, 30)
	splinter.AddComponent(engine.NewPhysics(true, 1, 1))
	splinter.Physics.Shape.IsSensor = true
	splinter.Tag = "splinter"
	for i := 0; i < 1; i++ {
		slc := splinter.Clone()
		slc.Transform().SetParent2(Layer3)
		slc.Transform().SetWorldPositionf(230, 130)
	}
	uvs, ind := engine.AnimatedGroupUVs(plAtlas, "player_walk", "player_stand", "player_attack", "player_jump", "player_bend", "player_hit", "player_climb")

	pl = engine.NewGameObject("Player")
	pl.AddComponent(engine.NewSprite3(plAtlas.Texture, uvs))
	pl.Sprite.BindAnimations(ind)
	pl.Sprite.AnimationSpeed = 10
	pl.Transform().SetWorldPositionf(100, 180)
	pl.Transform().SetWorldScalef(50, 100)
	pl.Transform().SetParent2(Layer2)
	pl.AddComponent(NewPlayer())
	pl.AddComponent(components.NewSmoothFollow(nil, 1, 30))
	pl.AddComponent(engine.NewPhysics(false, 1, 1))
	pl.Physics.Shape.SetFriction(0.7)
	pl.Physics.Shape.SetElasticity(0.2)
	pl.Tag = "player"

	uvs, ind = engine.AnimatedGroupUVs(enAtlas, "enemy_walk", "enemy_stand", "enemy_attack", "enemy_jump", "enemy_hit")
	en = engine.NewGameObject("Enemy")
	en.AddComponent(engine.NewSprite3(enAtlas.Texture, uvs))
	en.Sprite.BindAnimations(ind)
	en.Transform().SetWorldScalef(50, 100)
	en.AddComponent(engine.NewPhysics(false, 1, 1))
	en.Physics.Shape.SetFriction(0.7)
	en.Physics.Shape.SetElasticity(0.2)

	Hp := engine.NewGameObject("hpBar")
	Hp.GameObject().AddComponent(engine.NewSprite2(atlas.Texture, engine.IndexUV(atlas, spr_chudHp)))
	Hp.GameObject().Sprite.SetAlign(engine.AlignLeft)
	Hp.GameObject().Transform().SetWorldScalef(10, 15)

	for i := 0; i < 1; i++ {
		ec := en.Clone()
		Hp.GameObject().Transform().SetWorldPosition(engine.Vector{0, 580, 0})

		ec.Transform().SetWorldPositionf(200, 110)
		ec.AddComponent(NewEnemy((Hp.AddComponent(NewBar(17))).(*Bar)))
		ec.Transform().SetParent2(Layer2)

		ec.Sprite.AnimationSpeed = 10
		Hp.Transform().SetParent2(up)
	}
	box = engine.NewGameObject("box")
	box.AddComponent(engine.NewSprite2(atlas.Texture, engine.IndexUV(atlas, spr_box)))
	box.Transform().SetWorldScalef(40, 40)
	box.AddComponent(engine.NewPhysics(false, 1, 1))
	box.Physics.Body.SetMass(1.5)
	box.Physics.Shape.SetFriction(0.3)
	for i := 0; i < 1; i++ {
		bc := box.Clone()
		bc.Transform().SetParent2(Layer3)
		bc.Transform().SetWorldPositionf(30, 150)
	}

	lader = engine.NewGameObject("lader")
	lader.AddComponent(engine.NewSprite2(atlas.Texture, engine.IndexUV(atlas, spr_lader)))
	lader.Transform().SetWorldScalef(60, 100)
	lader.AddComponent(engine.NewPhysics(true, 1, 1))
	lader.Physics.Shape.IsSensor = true
	lader.Physics.Shape.SetFriction(2)
	lader.Tag = "lader"

	for i := 0; i < 1; i++ {
		lc := lader.Clone()
		lc.Transform().SetParent2(Layer3)
		lc.Transform().SetWorldPositionf(150, 150)
	}

	uvs, ind = engine.AnimatedGroupUVs(tileset, "ground")
	floor = engine.NewGameObject("floor")
	floor.AddComponent(engine.NewSprite3(tileset.Texture, uvs))
	floor.Sprite.BindAnimations(ind)
	floor.Sprite.AnimationSpeed = 0
	floor.Transform().SetWorldScalef(100, 100)
	floor.AddComponent(engine.NewPhysics(true, 1, 1))

	chud := engine.NewGameObject("chud")
	chud.AddComponent(engine.NewSprite2(atlas.Texture, engine.IndexUV(atlas, spr_chud)))
	ch = chud.AddComponent(NewChud()).(*Chud)
	chud.Transform().SetParent2(cam)
	chud.Transform().SetWorldPositionf(200, 550)
	chud.Transform().SetWorldScalef(100, 100)

	label := engine.NewGameObject("Label")
	label.Transform().SetParent2(cam)
	label.Transform().SetPositionf(20, float32(engine.Height)-40)
	label.Transform().SetScalef(20, 20)

	spot = engine.NewGameObject("spot")
	spot.AddComponent(engine.NewSprite2(atlas.Texture, engine.IndexUV(atlas, spr_spot)))
	// spot.Transform().SetParent2(Layer3)
	// spot.Transform().SetWorldPositionf(10, 150)
	spot.AddComponent(engine.NewPhysics(false, 1, 1))
	spot.Transform().SetScalef(30, 30)
	spot.AddComponent(NewItem(func() {
		plComp.AddHp(float32(rand.Int()%10 + 5))
		spot.Destroy()
	}))
	uvs, ind = engine.AnimatedGroupUVs(atlas, "chest")
	chest = engine.NewGameObject("chest")
	chest.AddComponent(engine.NewSprite3(atlas.Texture, uvs))
	chest.Sprite.BindAnimations(ind)
	chest.Transform().SetWorldScalef(70, 70)
	chest.AddComponent(engine.NewPhysics(false, 1, 1))
	chest.Physics.Shape.IsSensor = true
	chest.Sprite.AnimationSpeed = 0
	chest.Physics.Body.IgnoreGravity = true
	chest.AddComponent(NewChest(money))

	chest.Transform().SetWorldPositionf(300, 150)
	chest.Transform().SetParent2(Layer3)

	Hp = engine.NewGameObject("hpBar")
	Hp.GameObject().AddComponent(engine.NewSprite2(atlas.Texture, engine.IndexUV(atlas, spr_chudHp)))
	Hp.GameObject().Sprite.SetAlign(engine.AlignLeft)
	Hp.GameObject().Transform().SetWorldPosition(engine.Vector{235, 580, 0})
	Hp.GameObject().Transform().SetWorldScalef(17, 20)
	Hp.Transform().SetParent2(up)
	ch.Hp = (Hp.AddComponent(NewBar(17))).(*Bar)

	Cp := engine.NewGameObject("hpBar")
	Cp.GameObject().AddComponent(engine.NewSprite2(atlas.Texture, engine.IndexUV(atlas, spr_chudCp)))
	Cp.GameObject().Sprite.SetAlign(engine.AlignLeft)
	Cp.GameObject().Transform().SetWorldPosition(engine.Vector{235, 550, 0})
	Cp.GameObject().Transform().SetWorldScalef(17, 20)
	Cp.Transform().SetParent2(up)
	ch.Cp = (Cp.AddComponent(NewBar(17))).(*Bar)

	Exp := engine.NewGameObject("hpBar")
	Exp.GameObject().AddComponent(engine.NewSprite2(atlas.Texture, engine.IndexUV(atlas, spr_chudCp)))
	Exp.GameObject().Sprite.SetAlign(engine.AlignLeft)
	Exp.GameObject().Transform().SetWorldPosition(engine.Vector{235, 530, 0})
	Exp.GameObject().Transform().SetWorldScalef(17, 20)
	Exp.Transform().SetParent2(up)
	ch.Exp = (Exp.AddComponent(NewBar(17))).(*Bar)
	ch.Exp.Start()
	ch.Exp.SetValue(0, 100)
	money := engine.NewGameObject("money")
	money.Transform().SetParent2(cam)
	money.Transform().SetWorldPositionf(100, 500)
	money.Transform().SetScalef(20, 20)
	ch.Money = money.AddComponent(components.NewUIText(ArialFont2, "0")).(*components.UIText)
	ch.Money.SetAlign(engine.AlignLeft)
	txt2 := label.AddComponent(NewTestBox(func(tx *TestTextBox) {
		money.Transform().SetPositionf(235, float32(tx.V))
	})).(*TestTextBox)

	level := engine.NewGameObject("money")
	level.Transform().SetParent2(cam)
	level.Transform().SetWorldPositionf(50, 500)
	level.Transform().SetScalef(20, 20)
	ch.Level = level.AddComponent(components.NewUIText(ArialFont2, "1")).(*components.UIText)
	ch.Level.SetAlign(engine.AlignLeft)
	txt2.SetAlign(engine.AlignLeft)

	for i := 0; i < 10; i++ {
		f := floor.Clone()
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
	s.AddGameObject(Background)

}
func LoadTextures() {
	atlas = engine.NewManagedAtlas(2048, 1024)
	plAtlas = engine.NewManagedAtlas(2048, 1024)
	enAtlas = engine.NewManagedAtlas(2048, 1024)
	tileset = engine.NewManagedAtlas(2048, 1024)
	menuAtlas = engine.NewManagedAtlas(2048, 1024)
	CheckError(atlas.LoadImageID("./data/background/backGame.png", spr_bg))
	CheckError(atlas.LoadImageID("./data/objects/lader.png", spr_lader))
	CheckError(atlas.LoadImageID("./data/objects/splinter.png", spr_splinter))
	CheckError(atlas.LoadImageID("./data/objects/box.png", spr_box))
	CheckError(atlas.LoadImageID("./data/bar/chud.png", spr_chud))
	CheckError(atlas.LoadImageID("./data/bar/hpBar.png", spr_chudHp))
	CheckError(atlas.LoadImageID("./data/bar/cpBar.png", spr_chudCp))
	CheckError(atlas.LoadImageID("./data/bar/expBar.png", spr_chudExp))
	atlas.LoadGroupSheet("./data/items/chest.png", 41, 54, 4)
	CheckError(atlas.LoadImageID("./data/items/Coin.png", spr_coin))
	CheckError(atlas.LoadImageID("./data/items/Coin10.png", spr_coin10))
	CheckError(atlas.LoadImageID("./data/items/Daimond.png", spr_diamond))
	CheckError(atlas.LoadImageID("./data/items/spot.png", spr_spot))

	e, id := plAtlas.LoadGroupSheet("./data/player/player_walk.png", 187, 338, 4)
	CheckError(e)
	e, id = plAtlas.LoadGroupSheet("./data/player/player_stand.png", 187, 338, 1)
	CheckError(e)
	e, id = plAtlas.LoadGroupSheet("./data/player/player_attack.png", 249, 340, 9)
	CheckError(e)
	e, id = plAtlas.LoadGroupSheet("./data/player/player_jump.png", 236, 338, 1)
	CheckError(e)
	e, id = plAtlas.LoadGroupSheet("./data/player/player_bend.png", 188, 259, 1)
	CheckError(e)
	e, id = plAtlas.LoadGroupSheet("./data/player/player_climb.png", 236, 363, 2)
	CheckError(e)
	e, id = plAtlas.LoadGroupSheet("./data/player/player_hit.png", 206, 334, 1)
	CheckError(e)

	e, id = enAtlas.LoadGroupSheet("./data/Enemy/enemy_walk.png", 187, 338, 4)
	CheckError(e)
	e, id = enAtlas.LoadGroupSheet("./data/Enemy/enemy_stand.png", 187, 338, 1)
	CheckError(e)
	e, id = enAtlas.LoadGroupSheet("./data/Enemy/enemy_attack.png", 249, 340, 9)
	CheckError(e)
	e, id = enAtlas.LoadGroupSheet("./data/Enemy/enemy_jump.png", 236, 338, 1)
	CheckError(e)
	e, id = enAtlas.LoadGroupSheet("./data/Enemy/enemy_hit.png", 206, 334, 1)
	CheckError(e)

	e, id = tileset.LoadGroupSheet("./data/tileset/ground.png", 32, 32, 110)
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

	enAtlas.BuildAtlas()
	enAtlas.BuildMipmaps()
	enAtlas.SetFiltering(engine.MipMapLinearNearest, engine.Nearest)
	enAtlas.Texture.SetReadOnly()

	plAtlas.BuildAtlas()
	plAtlas.BuildMipmaps()
	plAtlas.SetFiltering(engine.MipMapLinearNearest, engine.Nearest)
	plAtlas.Texture.SetReadOnly()

	menuAtlas.BuildAtlas()
	menuAtlas.BuildMipmaps()
	menuAtlas.SetFiltering(engine.MipMapLinearNearest, engine.Nearest)
	menuAtlas.Texture.SetReadOnly()

	tileset.BuildAtlas()
	tileset.BuildMipmaps()
	tileset.SetFiltering(engine.MipMapLinearNearest, engine.Nearest)
	tileset.Texture.SetReadOnly()
}
func (s *PirateScene) New() engine.Scene {
	gs := new(PirateScene)
	gs.SceneData = engine.NewScene("PirateScene")
	return gs
}
