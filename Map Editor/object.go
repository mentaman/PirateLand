package main

import (
	"github.com/vova616/garageEngine/engine"

	"github.com/vova616/garageEngine/engine/input"
)

var (
	id      int       = 0
	objList []*Object = []*Object{}
)

func ClearList() {
	for _, v := range objList {
		v.GameObject().Destroy()
	}
	objList = []*Object{}
}

type Object struct {
	engine.BaseComponent
	mouseIn bool
	My_id   int
}

func (ob *Object) Start() {
	id++
	ob.My_id = id
	objList = append(objList, ob)
}
func (ob *Object) OnMouseEnter(a engine.Arbiter) bool {
	ob.mouseIn = true
	return false
}

func (ob *Object) OnMouseExit(a engine.Arbiter) {
	ob.mouseIn = false
}
func (ob *Object) Update() {
	if (ob.mouseIn && (input.MousePress(input.MouseRight) || (input.MouseDown(input.MouseRight) && input.KeyDown(input.KeyLshift)))) || (input.KeyDown(input.KeyLalt) && input.KeyDown('A')) {

		ob.GameObject().Destroy()
	}
}
func (ob *Object) OnDestroy() {
	for i, o := range objList {
		if o.My_id == ob.My_id {
			objList = append(objList[0:i], objList[i+1:]...)
			break
		}
	}
}
func NewObject() *Object {

	return &Object{engine.NewComponent(), false, 0}
}
