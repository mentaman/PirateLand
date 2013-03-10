package Background

import (
	"github.com/vova616/garageEngine/engine"
	"math"
)

var (
	Object *engine.GameObject
	Atlas  *engine.ManagedAtlas
)

const (
	Spr_bg = 1
)

type back struct {
	engine.BaseComponent
}

func NewBack() *back {
	return &back{engine.NewComponent()}
}
func Create() {
	Object = engine.NewGameObject("bg")
	Object.AddComponent(engine.NewSprite2(Atlas.Texture, engine.IndexUV(Atlas, Spr_bg)))
	Object.Transform().SetScalef(2000, 1800)
	Object.Transform().SetPositionf(-float32(engine.Width)/2, -float32(engine.Height)/2)
	Object.AddComponent(NewBack())
	Object.Transform().SetDepth(-2)
}
func (b *back) Update() {
	camera := engine.GetScene().SceneBase().Camera
	if camera != nil {
		myPos := engine.Vector{b.Transform().Position().X - float32(engine.Width/2), b.Transform().Position().Y - float32(engine.Height/2), 0}
		camPos := camera.Transform().Position()

		myPos = engine.Lerp(myPos, camPos, float32(engine.DeltaTime())*5)
		disX := myPos.X - camPos.X
		disY := myPos.Y - camPos.Y
		var MaxDis float32 = 300
		if float32(math.Abs(float64(disX))) > MaxDis {
			if disX < 0 {
				myPos.X = MaxDis - myPos.X
			} else {
				myPos.X = myPos.X + MaxDis
			}
		}
		if float32(math.Abs(float64(disY))) > MaxDis {
			if disY < 0 {
				myPos.Y = MaxDis - myPos.Y
			} else {
				myPos.Y = myPos.Y + MaxDis
			}
		}

		b.Transform().SetPosition(camPos)
	}
}
