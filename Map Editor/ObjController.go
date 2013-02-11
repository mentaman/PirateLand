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
	last   engine.Vector
}

func (m *ObjController) Start() {
	uvs, ind := engine.AnimatedGroupUVs(atlas, "player_walk", "chest", "enemy_walk")
	m.guiObj = engine.NewGameObject("guiObj")
	m.guiObj.AddComponent(engine.NewSprite3(atlas.Texture, uvs))
	m.guiObj.Sprite.BindAnimations(ind)
	m.guiObj.Sprite.AnimationSpeed = 0
	m.guiObj.Transform().SetScalef(m.width, m.height)
	m.guiObj.Transform().SetParent2(cam)

}
func (m *ObjController) Update() {

	px, py := input.MousePosition()
	if input.MousePress(input.MouseLeft) {
		cl := obj.Clone()
		cm := cam.Transform().Position()

		m.last = cm
		cl.Transform().SetPositionf(float32(px)+cm.X, float32(engine.Height-py)+(cm.Y))
		cl.Transform().SetScalef(m.width, m.height)
		cl.Transform().SetParent2(Layer1)
	}
	if input.KeyPress('L') {
		cam.Transform().SetPosition(m.last)
	}
	if input.KeyDown('W') {
		if input.KeyDown(input.KeyLshift) {
			m.width--
		} else {
			m.width++
		}
		m.guiObj.Transform().SetScalef(m.width, m.height)
	}
	if input.KeyDown('H') {
		if input.KeyDown(input.KeyLshift) {
			m.height--
		} else {
			m.height++
		}
		m.guiObj.Transform().SetScalef(m.width, m.height)
	}
	if input.KeyPress('S') {
		SaveXML()
	}
	m.guiObj.Transform().SetPositionf(float32(px), float32(engine.Height-py))

}

func NewObjController() *ObjController {
	return &ObjController{engine.NewComponent(), 30, 30, nil, engine.Vector{0, 0, 0}}
}
