package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Place struct {
	X float32 `xml:"Place>X"`
	Y float32 `xml:"Place>Y"`
	Z float32 `xml:"Place>Z"`
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

			// objM.Name = ob.GameObject().Sprite.CurrentAnimation().(string)
			objM.Index = int(ob.GameObject().Sprite.CurrentAnimationIndex())
			vP := ob.Transform().WorldPosition()
			vS := ob.Transform().WorldScale()
			objM.Iplace = Place{vP.X, vP.Y, vP.Z}
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
