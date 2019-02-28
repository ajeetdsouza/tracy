package main

import (
	"math"
	"os"

	"github.com/ajeetdsouza/tracy/display"
	"github.com/ajeetdsouza/tracy/geom"
)

func main() {
	size := 100

	canvas := display.NewCanvas(size, size)
	red := geom.NewColor(1, 0, 0)

	hand := geom.NewPoint(0, 3.0/16*float64(size), 0)

	for hour := 0; hour < 12; hour++ {
		rotation := float64(hour) / 6 * math.Pi
		translation := float64(size) / 2
		transform := geom.NewTransform().RotateZ(rotation).Translate(translation, translation, 0)

		currHand := hand.ApplyTransform(transform)
		x := int(math.Round(currHand.X()))
		y := int(math.Round(currHand.Y()))

		canvas.Grid[x*canvas.Height+y] = red
	}

	canvas.WritePpm(os.Stdout)
}
