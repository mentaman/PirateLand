package Game

import (
	"encoding/xml"
	"fmt"
	"github.com/mentaman/PirateLand/Game/Enemy"
	"github.com/mentaman/PirateLand/Game/GUI"
	"github.com/mentaman/PirateLand/Game/Objects"
	"github.com/mentaman/PirateLand/Game/Player"
	"io/ioutil"
)

type Place struct {
	X float32 `xml:"Place>X"`
	Y float32 `xml:"Place>Y"`
}
type Scale struct {
	Width  float32 `xml:"width"`
	Height float32 `xml:"height"`
}
type XObject struct {
	Name   string `xml:"name"`
	Index  int    `xml:"index"`
	Iplace Place  `xml:"place"`
	Iscale Scale  `xml:"scale"`
}
type XObjects struct {
	XMLName   xml.Name  `xml:"Objects"`
	Objs      []XObject `xml:"object"`
	CamPlace  Place
	LastPlace Place
}

func LoadMap(name string) {
	v := &XObjects{}
	cont, _ := ioutil.ReadFile(name)
	data := string(cont)
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
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
