package Game

import (
	"github.com/vova616/GarageEngine/Engine"
)

var (
	atlas *Engine.ManagedAtlas
)

type PirateScene struct {
	*Engine.Scenedata
}

func (s *PirateScene) SceneBase() *Engine.SceneData {
	return s.SceneData
}
func (s *PirateScene) Load() {
	LoadTexture()
}
func LoadTextures() {

}
func (s *PirateScene) New() Engine.Scene {
	gs := new(PirateScene)
	gs.SceneData = Engine.NewScene("PirateScene")
	return gs
}
