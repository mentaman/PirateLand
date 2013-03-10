package main

import (
	"github.com/vova616/garageEngine/engine"
	"github.com/vova616/garageEngine/engine/input"
)

type ObjController struct {
	engine.BaseComponent
	width    float32
	height   float32
	guiObj   *engine.GameObject
	Last     engine.Vector
	spriteId int
	grid     float32
}

func (m *ObjController) Start() {
	uvs, ind := engine.AnimatedGroupUVs(atlas, sprites...)
	m.guiObj = engine.NewGameObject("guiObj")
	m.guiObj.AddComponent(engine.NewSprite3(atlas.Texture, uvs))
	m.guiObj.AddComponent(engine.NewPhysics(false))
	m.guiObj.Sprite.BindAnimations(ind)
	m.guiObj.Sprite.AnimationSpeed = 0
	m.guiObj.Sprite.SetAlign(engine.AlignCenter)
	m.guiObj.Transform().SetScalef(m.width, m.height)
	m.guiObj.Transform().SetParent2(cam)
	m.guiObj.Sprite.SetAnimation(sprites[m.spriteId])
	m.guiObj.Tag = "guiObj"

}
func SnapToGrid(x, size float32) float32 {
	if size != 0 {
		return (size) * float32(int((x / size)))
	}
	return x
}
func (m *ObjController) Update() {

	px, py := input.MousePosition()

	m.guiObj.Transform().SetPositionf(SnapToGrid(float32(px)+m.grid/2, m.grid)+m.grid/2-640, SnapToGrid(float32(engine.Height-py)+m.grid/2, m.grid)+m.grid/2-360)
	guiP := m.guiObj.Transform().WorldPosition()
	if input.MousePress(input.MouseLeft) {
		cl := obj.Clone()
		cm := cam.Transform().Position()

		m.Last = cm
		cl.Transform().SetPositionf(guiP.X, guiP.Y)
		cl.Transform().SetScalef(m.width, m.height)
		cl.Transform().SetParent2(Layer1)

		cl.Sprite.SetAnimation(m.guiObj.Sprite.CurrentAnimation().(string))
		cl.Sprite.SetAnimationIndex(int(m.guiObj.Sprite.CurrentAnimationIndex()))
	}
	if input.KeyPress('I') {
		if input.KeyDown(input.KeyLshift) {
			m.spriteId--
			if m.spriteId < 0 {
				m.spriteId = len(sprites) - 1
			}
		} else {
			m.spriteId++
			if m.spriteId > len(sprites)-1 {
				m.spriteId = 0
			}
		}

		m.guiObj.Sprite.SetAnimation(sprites[m.spriteId])

	}
	if input.KeyPress('L') {
		cam.Transform().SetPosition(m.Last)
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
		SaveGOB()
	}
	if input.KeyPress('R') {
		LoadGOB()
	}
	if input.KeyPress('C') {
		ClearList()
	}
	if input.KeyPress(input.Key_Right) {
		dd := m.guiObj
		dd.GameObject().Sprite.SetAnimationIndex((int(dd.GameObject().Sprite.CurrentAnimationIndex()) + 1) % dd.GameObject().Sprite.AnimationLength())
	} else if input.KeyPress(input.Key_Left) {
		dd := m.guiObj
		b := (int(dd.GameObject().Sprite.CurrentAnimationIndex()) - 1)

		if b < 0 {
			b = dd.GameObject().Sprite.AnimationLength() - 1
		} else {
			b = b % (dd.GameObject().Sprite.AnimationLength())
		}
		dd.GameObject().Sprite.SetAnimationIndex(b)
	}
	if input.KeyDown('G') {
		if input.KeyDown(input.KeyLshift) {
			m.grid--
			if m.grid < 0 {
				m.grid = 0
			}
		} else {
			m.grid++
		}
	}

}

func NewObjController() *ObjController {
	return &ObjController{engine.NewComponent(), 60, 60, nil, engine.Vector{0, 0, 0}, 0, 30}
}
