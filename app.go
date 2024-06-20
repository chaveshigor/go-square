package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"time"

	"github.com/chaveshigor/square-go/shapes"
	"github.com/chaveshigor/square-go/solids"
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

// DrawPolygon draws lines between the given points
func DrawPolygon(img *image.RGBA, points []shapes.Point2d, col color.Color) {
	if len(points) < 2 {
		return
	}
	for i := 0; i < len(points)-1; i++ {
		currentPoint := image.Point{int(points[i][0]), int(points[i][1])}
		nextPoint := image.Point{int(points[i+1][0]), int(points[i+1][1])}
		drawLine(img, currentPoint, nextPoint, col)
	}
}

// drawLine draws a line from point p1 to point p2
func drawLine(img *image.RGBA, p1, p2 image.Point, col color.Color) {
	dx := abs(p2.X - p1.X)
	dy := abs(p2.Y - p1.Y)
	sx, sy := 1, 1
	if p1.X >= p2.X {
		sx = -1
	}
	if p1.Y >= p2.Y {
		sy = -1
	}
	err := dx - dy
	for {
		img.Set(p1.X, p1.Y, col)
		if p1 == p2 {
			break
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			p1.X += sx
		}
		if e2 < dx {
			err += dx
			p1.Y += sy
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(&screen.NewWindowOptions{
			Width:  1280,
			Height: 768,
			Title:  "Draw Polygon",
		})
		if err != nil {
			log.Fatal(err)
		}
		defer w.Release()

		screenBuffer, err := s.NewBuffer(image.Point{640, 480})
		if err != nil {
			log.Fatal(err)
		}
		defer screenBuffer.Release()

		square := solids.NewSquare(
			solids.Point3d{X: 300, Y: 100, Z: 100},
			solids.Point3d{X: 400, Y: 100, Z: 100},
			solids.Point3d{X: 400, Y: 100, Z: 200},
			solids.Point3d{X: 300, Y: 100, Z: 200},
			solids.Point3d{X: 300, Y: 100, Z: 100},
			solids.Point3d{X: 300, Y: 200, Z: 100},
			solids.Point3d{X: 400, Y: 200, Z: 100},
			solids.Point3d{X: 400, Y: 100, Z: 100},
			solids.Point3d{X: 400, Y: 200, Z: 100},
			solids.Point3d{X: 400, Y: 200, Z: 200},
			solids.Point3d{X: 400, Y: 100, Z: 200},
			solids.Point3d{X: 400, Y: 200, Z: 200},
			solids.Point3d{X: 300, Y: 200, Z: 200},
			solids.Point3d{X: 300, Y: 100, Z: 200},
			solids.Point3d{X: 300, Y: 200, Z: 200},
			solids.Point3d{X: 300, Y: 200, Z: 100},
		)

		angle := 0
		for {
			e := w.NextEvent()
			switch e := e.(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			case paint.Event:
				img := screenBuffer.RGBA()

				// Clear the screen with black color
				for {
					rotatedSquare := solids.RotateSolid(square, float64(angle))
					shape := shapes.Transpose(rotatedSquare)
					points := shape.Points

					draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.Point{}, draw.Src)

					// Draw the polygon
					DrawPolygon(img, points, color.White)

					w.Upload(image.Point{}, screenBuffer, screenBuffer.Bounds())
					w.Publish()

					time.Sleep(44 * time.Millisecond)
					angle = angle + 1
					fmt.Println(angle)
				}
			case size.Event:
				// Handle resizing here if needed
			}
		}
	})
}
