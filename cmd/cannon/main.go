package main

import (
	"math"
	"os"

	"github.com/ajeetdsouza/tracy/display"
	"github.com/ajeetdsouza/tracy/geom"
)

type Environment struct {
	grav, wind geom.Tuple
}

type Projectile struct {
	pos, vel geom.Tuple
}

func (proj *Projectile) tick(env Environment) {
	proj.pos = proj.pos.Add(proj.vel)
	proj.vel = proj.vel.Add(env.grav).Add(env.wind)
}

func main() {
	proj := Projectile{
		pos: geom.NewPoint(0, 1, 0),
		vel: geom.NewVector(1, 1.8, 0).Normalize().Mul(11.25),
	}

	env := Environment{
		grav: geom.NewVector(0, -0.1, 0),
		wind: geom.NewVector(-0.01, 0, 0),
	}

	display := display.NewCanvas(900, 550)
	red := geom.NewColor(1, 0, 0)

	for proj.pos.Y() >= 0 {
		x := int(math.Round(proj.pos.X()))
		y := int(math.Round(proj.pos.Y()))

		if x < display.Width && y < display.Height {
			display.Grid[x*display.Height+y] = red
		}

		proj.tick(env)
	}

	display.WritePpm(os.Stdout)
}
