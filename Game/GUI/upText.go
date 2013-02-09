package GUI

import (
	"github.com/mentaman/PirateLand/Game/Fonts"
	"github.com/vova616/garageEngine/engine"
	"github.com/vova616/garageEngine/engine/components"
	"github.com/vova616/garageEngine/engine/components/tween"
	"time"
)

type upText struct {
	engine.BaseComponent
	label   *engine.GameObject
	text    *components.UIText
	content string
}

func NewUpText(c string) *upText {
	return &upText{engine.NewComponent(), nil, nil, c}
}
func (u *upText) Start() {
	label := engine.NewGameObject("Label")
	label.Transform().SetParent(u.Transform().Parent())
	p := u.Transform().WorldPosition()
	si := u.Transform().WorldScale()
	label.Transform().SetWorldPositionf(p.X, p.Y)
	label.Transform().SetScalef(si.X, si.Y)
	txt2 := label.AddComponent(components.NewUIText(Fonts.ArialFont2, u.content)).(*components.UIText)
	u.text = txt2
	u.label = label
	tween.Create(&tween.Tween{Target: label.GameObject(), From: []float32{p.Y}, To: []float32{p.Y + 100},
		Algo: tween.Linear, Type: tween.WorldPosition, Time: time.Second * 4, Loop: tween.None, Format: "y", EndCallback: func() { u.label.GameObject().Destroy(); u.GameObject().Destroy() }})
}
func NewUpTextObj(c string, t *engine.Transform, size float32) {
	te := engine.NewGameObject("textup")
	te.Transform().SetParent(t.Parent())
	te.Transform().SetPosition(t.Position())
	te.Transform().SetScalef(size, size)
	te.AddComponent(NewUpText(c))
}
