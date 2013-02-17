package main

import (
	// "encoding/xml"
	"encoding/gob"
	"fmt"
	"io/ioutil"

	"bytes"
	"os"
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

func SaveGOB() {

	v := &XObjects{}
	v.Objs = []XObject{}
	c := cam.Transform().Position()
	v.CamPlace = Place{c.X, c.Y}
	cl := objControll.Last
	v.LastPlace = Place{cl.X, cl.Y}
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
	m := new(bytes.Buffer)
	enc := gob.NewEncoder(m)
	err := enc.Encode(v)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	file, _ := os.Create("./map.txt")
	defer file.Close()
	file.Write(m.Bytes())
}
func LoadGOB() {

	cont, _ := ioutil.ReadFile("map.txt")
	v := &XObjects{}
	p := bytes.NewBuffer(cont)
	dec := gob.NewDecoder(p)
	err := dec.Decode(&v)
	if err != nil {
		panic(err)
	}
	ClearList()
	cam.Transform().SetPositionf(v.CamPlace.X, v.CamPlace.Y)
	objControll.Transform().SetPositionf(v.LastPlace.X, v.LastPlace.Y)
	for _, robj := range v.Objs {
		cl := obj.Clone()

		cl.Transform().SetWorldPositionf(robj.Iplace.X, robj.Iplace.Y)
		cl.Transform().SetWorldScalef(robj.Iscale.Width, robj.Iscale.Height)
		cl.Transform().SetParent2(Layer1)

		cl.Sprite.SetAnimation(robj.Name)
		cl.Sprite.SetAnimationIndex(robj.Index)
	}
}
