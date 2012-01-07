/*
* Copyright (c) 2011 Matteo Goggi
*
* This software is provided 'as-is', without any express or implied 
* warranty.  In no event will the authors be held liable for any damages 
* arising from the use of this software. 
* Permission is granted to anyone to use this software for any purpose, 
* including commercial applications, and to alter it and redistribute it 
* freely, subject to the following restrictions: 
* 1. The origin of this software must not be misrepresented; you must not 
* claim that you wrote the original software. If you use this software 
* in a product, an acknowledgment in the product documentation would be 
* appreciated but is not required. 
* 2. Altered source versions must be plainly marked as such, and must not be 
* misrepresented as being the original software. 
* 3. This notice may not be removed or altered from any source distribution. 
*/
package mater

import (
	"math"
	"gl"
	"mater/vect"
)

type Camera struct {
	ScreenSize, Position, Scale vect.Vect
	Rotation float64
}

func (cam Camera) PreDraw() {
	gl.PushMatrix()
	gl.Translated(cam.ScreenSize.X / 2, cam.ScreenSize.Y / 2, 0)

	gl.Rotated(360 - cam.Rotation, 0, 0, 1)
	gl.Scaled(cam.Scale.X, cam.Scale.Y, 1)

	gl.Translated(-cam.Position.X, -cam.Position.Y, 0)
}

func (cam Camera) PostDraw() {
	gl.LoadIdentity()
	gl.PopMatrix()
}

func (cam Camera) Move(delta vect.Vect) {
	cam.Position = vect.Add(cam.Position, delta)
}

func (cam Camera) WorldToScreen (worldPos vect.Vect) vect.Vect {
	c := math.Cos(-cam.Rotation)
	s := math.Sin(-cam.Rotation)
	
	tx := worldPos.X - cam.Position.X
	tx += (cam.ScreenSize.X / 2) * cam.Scale.X
	
	ty := worldPos.Y - cam.Position.Y
	ty += (cam.ScreenSize.Y / 2) * cam.Scale.Y
	
	sx := c * tx - s * ty
	sy := s * tx + c * ty
	
	return vect.Vect{sx * cam.Scale.X, sy * cam.Scale.Y}
}

func (cam Camera) ScreenToWorld (screenPos vect.Vect) vect.Vect {
	c := math.Cos(cam.Rotation)
	s := math.Sin(cam.Rotation)
	
	tx := screenPos.X / cam.Scale.X
	ty := screenPos.Y / cam.Scale.Y
	
	sx := c * tx - s * ty
	sy := s * tx + c * ty
	
	sx += cam.Position.X
	sy += cam.Position.Y
	
	sx -= (cam.ScreenSize.X / 2) / cam.Scale.X
	sy -= (cam.ScreenSize.Y / 2) / cam.Scale.Y
	
	return vect.Vect{sx, sy}
}


