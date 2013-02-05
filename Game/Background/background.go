package Background

import (
	"github.com/vova616/garageEngine/engine"
)

var (
	Object *engine.GameObject
	Atlas  *engine.ManagedAtlas
)

const (
	Spr_bg = 1
)

func Create() {
	Object = engine.NewGameObject("bg")
	Object.AddComponent(engine.NewSprite2(Atlas.Texture, engine.IndexUV(Atlas, Spr_bg)))
	Object.Transform().SetScalef(2000, 1800)
	Object.Transform().SetPositionf(0, 0)
}
