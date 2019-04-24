package main

import (
	"math"
	"os"

	"github.com/ajeetdsouza/tracy/pkg/display"
	"github.com/ajeetdsouza/tracy/pkg/engine"
)

type Environment struct {
	grav, wind engine.Tuple
}

type Projectile struct {
	pos, vel engine.Tuple
}

func (proj *Projectile) tick(env Environment) {
	proj.pos = proj.pos.Add(proj.vel)
	proj.vel = proj.vel.Add(env.grav).Add(env.wind)
}

func main() {
	proj := Projectile{
		pos: engine.NewPoint(0, 1, 0),
		vel: engine.NewVector(1, 1.8, 0).Normalize().Mul(11.25),
	}

	env := Environment{
		grav: engine.NewVector(0, -0.1, 0),
		wind: engine.NewVector(-0.01, 0, 0),
	}

	canvas := display.NewCanvas(900, 550)
	red := engine.NewColor(1, 0, 0)

	for proj.pos.Y() >= 0 {
		x := int(math.Round(proj.pos.X()))
		y := canvas.Height - int(math.Round(proj.pos.Y()))

		if 0 <= x && x < canvas.Width && 0 <= y && y < canvas.Height {
			canvas.Grid[x*canvas.Height+y] = red
		}

		proj.tick(env)
	}

	canvas.WritePpm(os.Stdout)
}
