package main

import (
	"github.com/vova616/garageEngine/engine"
	"github.com/vova616/garageEngine/engine/input"
)

type ObjController struct {
	engine.BaseComponent
	width  float32
	height float32
	guiObj *engine.GameObject
}

func (m *ObjController) Start() {
	uvs, ind := engine.AnimatedGroupUVs(atlas, "player_walk", "chest", "enemy_walk")
	m.guiObj = engine.NewGameObject("guiObj")
	m.guiObj.AddComponent(engine.NewSprite3(atlas.Texture, uvs))
	m.guiObj.Sprite.BindAnimations(ind)
	m.guiObj.Sprite.AnimationSpeed = 0
}
func (m *ObjController) Update() {
	if input.MousePress(input.MouseLeft) {
		cl := obj.Clone()
		px, py := input.MousePosition()
		cm := cam.Transform().Position()
		cl.Transform().SetPositionf(float32(px)+cm.X, float32(engine.Height-py)+(cm.Y))
		cl.Transform().SetScalef(m.width, m.height)
		cl.Transform().SetParent2(Layer1)
	}
	if input.KeyDown('W') {
		m.width++
	}
	if input.KeyDown('H') {
		m.height++
	}
}

func NewObjController() *ObjController {
	return &ObjController{engine.NewComponent(), 30, 30}
}
