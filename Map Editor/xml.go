package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
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
	XMLName xml.Name  `xml:"Objects"`
	Objs    []XObject `xml:"object"`
}

func SaveXML() {
	v := &XObjects{}
	v.Objs = []XObject{}
	for _, ob := range objList {
		objM := XObject{}
		if ob.GameObject() != nil {
			objM.Name = ob.GameObject().Sprite.CurrentAnimation().(string)
			objM.Index = int(ob.GameObject().Sprite.CurrentAnimationIndex())
			vP := ob.Transform().WorldPosition()
			vS := ob.Transform().WorldScale()
			objM.Iplace = Place{vP.X, vP.Y}
			objM.Iscale = Scale{vS.X, vS.Y}
			v.Objs = append(v.Objs, objM)
		}

	}

	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	file, _ := os.Create("./map.txt")
	defer file.Close()
	file.Write(output)
}
func LoadXML() {
	v := &XObjects{}
	cont, _ := ioutil.ReadFile("map.txt")
	data := string(cont)
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	ClearList()
	for _, robj := range v.Objs {
		cl := obj.Clone()

		cl.Transform().SetWorldPositionf(robj.Iplace.X, robj.Iplace.Y)
		cl.Transform().SetWorldScalef(robj.Iscale.Width, robj.Iscale.Height)
		cl.Transform().SetParent2(Layer1)

		cl.Sprite.SetAnimation(robj.Name)
		cl.Sprite.SetAnimationIndex(robj.Index)
	}
}
