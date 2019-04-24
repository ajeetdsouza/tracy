package display

import (
	"fmt"
	"io"
	"math"

	"github.com/ajeetdsouza/tracy/pkg/engine"
)

type Canvas struct {
	Width, Height int
	Grid          []engine.Color
}

func NewCanvas(width, height int) Canvas {
	pixels := width * height
	canvas := Canvas{width, height, make([]engine.Color, pixels)}
	for i := 0; i < pixels; i++ {
		canvas.Grid[i] = engine.NewColor(0, 0, 0)
	}
	return canvas
}

func (canvas Canvas) WritePpm(writer io.Writer) {
	fmt.Fprintf(writer, "P3\n%d %d\n255\n", canvas.Width, canvas.Height)
	for row := 0; row < canvas.Height; row++ {
		for col := 0; col < canvas.Width; col++ {
			color := canvas.Grid[col*canvas.Height+row]
			ir := scaleTo255(color.R())
			ig := scaleTo255(color.G())
			ib := scaleTo255(color.B())
			fmt.Fprintf(writer, "%d %d %d\n", ir, ig, ib)
		}
	}
}

func scaleTo255(val float64) int {
	val = math.Round(val * 255)
	val = math.Max(0, val)
	val = math.Min(255, val)
	return int(val)
}
