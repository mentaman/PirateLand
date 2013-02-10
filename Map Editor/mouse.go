package main

import (
	"github.com/vova616/garageEngine/engine"
)

type MouseController struct {
	engine.BaseComponent
}

func (m MouseController) Update() {
	p := m.Transform().Position()
	cam := engine.GetScene().SceneBase().Camera
	// camPos = cam.Transform().Position()

	if p.X < 50 {

		np := cam.Transform().Position()
		np = np.Add(engine.NewVector2(-10, 0))
		cam.Transform().SetPosition(np)
	} else if p.X > float32(engine.Width)-50 {
		np := cam.Transform().Position()
		np = np.Add(engine.NewVector2(10, 0))
		cam.Transform().SetPosition(np)
	}
	if p.Y < 50 {
		np := cam.Transform().Position()
		np = np.Add(engine.NewVector2(0, -10))
		cam.Transform().SetPosition(np)
	}
	if p.Y > float32(engine.Height)-50 {
		np := cam.Transform().Position()
		np = np.Add(engine.NewVector2(0, 10))
		cam.Transform().SetPosition(np)
	}
}

func NewMouseController() *MouseController {
	return &MouseController{engine.NewComponent()}
}
