package main

import (
	"github.com/teomat/mater/camera"
	"github.com/teomat/mater/engine"
)

//called whenever a new component is added to an entity in the active scene
func onNewComponent(entity *engine.Entity, comp engine.Component) {
	if comp.Name() == "Camera" {
		cam := comp.(*camera.Camera)

		if cam.IsMainCamera {
			MainCamera = cam
			MainCamera.ScreenSize = ScreenSize
		}
	}
}
