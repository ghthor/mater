package render

import (
	gl "github.com/banthar/Go-OpenGL/gl"
	"github.com/teomat/mater/vect"
	"math"
)

func DrawQuad(upperLeft, lowerRight vect.Vect, filled bool) {
	if filled {
		gl.Begin(gl.QUADS)
	} else {
		gl.Begin(gl.LINE_LOOP)
	}
	defer gl.End()

	gl.Vertex2d(upperLeft.X, upperLeft.Y)
	gl.Vertex2d(upperLeft.X, lowerRight.Y)
	gl.Vertex2d(lowerRight.X, lowerRight.Y)
	gl.Vertex2d(lowerRight.X, upperLeft.Y)
}

const (
	circlestep = 45
	deg2grad   = math.Pi / 180
)

func DrawCircle(pos vect.Vect, radius float64, filled bool) {
	if filled {
		gl.Begin(gl.TRIANGLE_FAN)
		gl.Vertex2d(pos.X, pos.Y)
	} else {
		gl.Begin(gl.LINE_LOOP)
	}
	defer gl.End()

	var d float64
	for i := 0.0; i <= 360; i += circlestep {
		d = deg2grad * i
		gl.Vertex2d(pos.X+math.Cos(d)*radius, pos.Y+math.Sin(d)*radius)
	}
}

func DrawLine(start, end vect.Vect) {
	gl.Begin(gl.LINES)
	defer gl.End()

	gl.Vertex2d(start.X, start.Y)
	gl.Vertex2d(end.X, end.Y)
}

func DrawPoly(vertices []vect.Vect, vertCount int, filled bool) {
	if filled {
		gl.Begin(gl.TRIANGLE_FAN)
		gl.Vertex2d(vertices[0].X, vertices[0].Y)
	} else {
		gl.Begin(gl.LINE_LOOP)
	}
	defer gl.End()

	for i := 0; i < vertCount; i++ {
		v := vertices[i]
		gl.Vertex2d(v.X, v.Y)
	}
}
