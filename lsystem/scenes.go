// scenes.go
package lsystem

import (
	"time"

	"github.com/donomii/sceneCamera"
	"github.com/donomii/wasm-rotating-cube/tween"
)

type Scene struct {
	Gallery   []string
	Selection int
	Active    bool
	Init      func(s *Scene)
	Camera    *sceneCamera.SceneCamera
	Clock     float32
}

func start_clock(v *float32, start, end, duration float32) {
	go tween.Clock(v, start, end, duration, time.Now())
}

func InitScenes(camera *sceneCamera.SceneCamera) []*Scene {
	var plantScene Scene
	plantScene.Camera = camera
	plantScene.Gallery = plantGallery
	plantScene.Init = func(s *Scene) {
		start_clock(&s.Clock, 0.0, 1000.0, 4000.0)
	}

	var renderScene Scene
	renderScene.Camera = camera
	renderScene.Gallery = renderGallery
	renderScene.Init = func(s *Scene) {
		start_clock(&s.Clock, 0.0, 1000.0, 12000.0)
	}

	var snekScene Scene
	snekScene.Camera = camera
	snekScene.Gallery = snekGallery
	snekScene.Init = func(s *Scene) {
		go transform_snek(s)
		start_clock(&s.Clock, 0.0, 1000.0, 4000.0)
	}

	return []*Scene{&plantScene, &renderScene, &snekScene}
}
