package Game

import (
	"bytes"
	"encoding/gob"
	"github.com/mentaman/PirateLand/Game/Enemy"
	"github.com/mentaman/PirateLand/Game/GUI"
	"github.com/mentaman/PirateLand/Game/Objects"
	"github.com/mentaman/PirateLand/Game/Player"
	"io/ioutil"
)

type Place struct {
	X float32
	Y float32
}
type Scale struct {
	Width  float32
	Height float32
}
type XObject struct {
	Name   string
	Index  int
	Iplace Place
	Iscale Scale
}
type XObjects struct {
	Objs      []XObject
	CamPlace  Place
	LastPlace Place
}

func LoadMap(name string) {
	cont, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	v := &XObjects{}
	p := bytes.NewBuffer(cont)
	dec := gob.NewDecoder(p)
	err = dec.Decode(&v)
	if err != nil {
		panic(err)
	}

	sd := GUI.NewBar(10)
	for _, robj := range v.Objs {
		switch robj.Name {
		case "player_walk":
			Player.Pl.Transform().SetPositionf(robj.Iplace.X, robj.Iplace.Y)
		case "ground":
			f := Objects.Floor.Clone()
			f.Sprite.SetAnimation("ground")

			f.Transform().SetWorldScalef(robj.Iscale.Width, robj.Iscale.Height)
			f.Transform().SetWorldPositionf(robj.Iplace.X, robj.Iplace.Y)
			f.Sprite.SetAnimationIndex(robj.Index)
			f.Transform().SetParent2(Layer3)
		case "splinter":
			slc := Objects.Splinter.Clone()
			slc.Transform().SetParent2(Layer3)
			slc.Transform().SetWorldPositionf(robj.Iplace.X, robj.Iplace.Y)
			slc.Transform().SetWorldScalef(robj.Iscale.Width, robj.Iscale.Height)
		case "enemy_walk":
			ec := Enemy.Regular.Clone()
			hpB := Enemy.HpBar.Clone()
			hpBd := hpB.ComponentTypeOfi(sd).(*GUI.Bar)

			ec.Transform().SetWorldPositionf(robj.Iplace.X, robj.Iplace.Y)
			ec.AddComponent(Enemy.NewEnemy(hpBd))
			ec.Transform().SetParent2(Layer2)

			ec.Sprite.AnimationSpeed = 10
			hpB.Transform().SetParent2(CamLayer)
		case "chest":
			s := Objects.ChestO.Clone()
			s.Transform().SetWorldPositionf(robj.Iplace.X, robj.Iplace.Y)
			s.Transform().SetParent2(Layer3)
		case "lader":
			lc := Objects.Lader.Clone()
			lc.Transform().SetParent2(Layer3)
			lc.Transform().SetWorldPositionf(robj.Iplace.X, robj.Iplace.Y)
		case "box":
			bc := Objects.Box.Clone()
			bc.Transform().SetParent2(Layer3)
			bc.Transform().SetWorldPositionf(robj.Iplace.X, robj.Iplace.Y)
		}

		// cl.Transform().SetWorldPositionf(robj.Iplace.X, robj.Iplace.Y)
		// cl.Transform().SetWorldScalef(robj.Iscale.Width, robj.Iscale.Height)
		// cl.Transform().SetParent2(Layer1)

		// cl.Sprite.SetAnimation(robj.Name)
		// cl.Sprite.SetAnimationIndex(robj.Index)
	}
}
