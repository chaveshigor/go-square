package shapes

import (
	"github.com/chaveshigor/square-go/solids"
)

type Point2d [2]float64

type Shape struct {
	Points []Point2d
}

var l float64 = 100

func Transpose(solid solids.Solid) Shape {
	var shape Shape
	var points []Point2d

	for _, v := range solid.Points {
		var d1x float64
		var d1y float64
		var d2 float64

		d1y = v.Y
		d1x = v.X
		d2 = v.Z

		var psx float64
		var psy float64

		psx = (d1x / d2) * l
		psy = (d1y / d2) * l

		psx = d1x
		psy = d1y

		newPoint := Point2d{psx, psy}
		points = append(points, newPoint)
	}

	shape.Points = points

	return shape
}
