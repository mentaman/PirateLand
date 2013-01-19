package Game

import (
	"fmt"
	"github.com/vova616/GarageEngine/Engine"
)

var (
	atlas            *Engine.ManagedAtlas
	GameSceneGeneral *PirateScene
)

type PirateScene struct {
	*Engine.SceneData
}

func (s *PirateScene) SceneBase() *Engine.SceneData {
	return s.SceneData
}
func (s *PirateScene) Load() {
	s.LoadTextures()
}
func (s *PirateScene) LoadTextures() {

}
func (s *PirateScene) New() Engine.Scene {
	gs := new(PirateScene)
	gs.SceneData = Engine.NewScene("PirateScene")
	return gs
}
