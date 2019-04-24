package main

import (
	"os"

	"github.com/ajeetdsouza/tracy/pkg/display"
	"github.com/ajeetdsouza/tracy/pkg/engine"
)

func main() {
	rayOrigin := engine.NewPoint(0, 0, -5)
	wallZ := 10.0

	wallSize := 7.0
	canvasSize := 100
	pixelSize := wallSize / float64(canvasSize)

	wallHalf := wallSize / 2.0

	canvas := display.NewCanvas(canvasSize, canvasSize)

	sphere := engine.Sphere{
		Material: engine.Material{
			Color:     engine.NewColor(1, 0.2, 1),
			Ambient:   0.1,
			Diffuse:   0.9,
			Specular:  0.9,
			Shininess: 200.0,
		},
		Transform: engine.NewTransform(),
	}

	light := engine.PointLight{
		Position:  engine.NewPoint(-10, 10, -10),
		Intensity: engine.NewColor(1, 1, 1),
	}

	// for each row of pixels in the canvas
	for y := 0; y < canvas.Height; y++ {
		worldY := wallHalf - pixelSize*float64(y)

		// for each pixel in the row
		for x := 0; x < canvas.Width; x++ {
			// compute the world x coordinate (left = -half, right = half)
			worldX := -wallHalf + pixelSize*float64(x)

			// describe the point on the wall that the ray will target
			position := engine.NewPoint(worldX, worldY, wallZ)

			ray := engine.Ray{
				Origin:    position,
				Direction: position.Sub(rayOrigin).Normalize(),
			}
			intersections := sphere.Intersect(ray)
			hit := intersections.Hit()

			if hit != nil {
				point := ray.At(hit.T)
				normal := hit.Shape.NormalAt(point)
				eye := ray.Direction.Neg()
				color := hit.Shape.GetMaterial().Lighting(light, point, eye, normal)

				canvas.Grid[x*canvas.Height+y] = color
			}
		}
	}

	canvas.WritePpm(os.Stdout)
}
