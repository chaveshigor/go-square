package solids

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type Point3d struct {
	X float64
	Y float64
	Z float64
}

type Solid struct {
	Points []Point3d
}

func NewSquare(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15, p16 Point3d) Solid {
	points := []Point3d{
		p1,
		p2,
		p3,
		p4,
		p5,
		p6,
		p7,
		p8,
		p9,
		p10,
		p11,
		p12,
		p13,
		p14,
		p15,
		p16,
	}

	return Solid{points}
}

func RotateSolid(solid Solid, angle float64) Solid {
	angleInRad := angle * math.Pi / 180

	var rotatedPoints []Point3d
	points := solid.Points
	for _, v := range points {
		pointsMatrix := mat.NewDense(3, 1, []float64{
			v.X - 350,
			v.Y - 150,
			v.Z - 150,
		})

		rotationMatrixY := mat.NewDense(3, 3, []float64{
			math.Cos(angleInRad), 0, math.Sin(angleInRad),
			0, 1, 0,
			-1 * math.Sin(angleInRad), 0, math.Cos(angleInRad),
		})

		rotationMatrixX := mat.NewDense(3, 3, []float64{
			1, 0, 0,
			0, math.Cos(angleInRad), -1 * math.Sin(angleInRad),
			0, math.Sin(angleInRad), math.Cos(angleInRad),
		})

		result := mat.NewDense(3, 1, nil)

		// fmt.Println(mat.Formatted(pointsMatrix))
		// fmt.Println(mat.Formatted(rotationMatrix))
		result.Mul(rotationMatrixY, pointsMatrix)
		result.Mul(rotationMatrixX, result)

		rotatedPoints = append(
			rotatedPoints,
			Point3d{X: result.At(0, 0) + 350, Y: result.At(1, 0) + 150, Z: result.At(2, 0) + 150},
		)

	}
	return Solid{Points: rotatedPoints}
}
