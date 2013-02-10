package main

import (
	"github.com/vova616/garageEngine/engine"

	"github.com/vova616/garageEngine/engine/input"
)

type Object struct {
	engine.BaseComponent
	mouseIn bool
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
func NewObject() *Object {
	return &Object{engine.NewComponent(), false}
}
