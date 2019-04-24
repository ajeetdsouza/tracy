package engine

import (
	"math"
	"testing"
)

func TestLighting(t *testing.T) {
	material := Material{
		Color:     NewColor(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
	position := NewPoint(0, 0, 0)

	tests := []struct {
		eye, normal Tuple
		light       PointLight
		result      Color
	}{
		{
			eye:    NewVector(0, 0, -1),
			normal: NewVector(0, 0, -1),
			light: PointLight{
				Position:  NewPoint(0, 0, -10),
				Intensity: NewColor(1, 1, 1),
			},
			result: NewColor(1.9, 1.9, 1.9),
		},
		{
			eye:    NewVector(0, 1/math.Sqrt2, -1/math.Sqrt2),
			normal: NewVector(0, 0, -1),
			light: PointLight{
				Position:  NewPoint(0, 0, -10),
				Intensity: NewColor(1, 1, 1),
			},
			result: NewColor(1.0, 1.0, 1.0),
		},
		{
			eye:    NewVector(0, 0, -1),
			normal: NewVector(0, 0, -1),
			light: PointLight{
				Position:  NewPoint(0, 10, -10),
				Intensity: NewColor(1, 1, 1),
			},
			result: NewColor(0.7364, 0.7364, 0.7364),
		},
		{
			eye:    NewVector(0, -1/math.Sqrt2, -1/math.Sqrt2),
			normal: NewVector(0, 0, -1),
			light: PointLight{
				Position:  NewPoint(0, 10, -10),
				Intensity: NewColor(1, 1, 1),
			},
			result: NewColor(1.6364, 1.6364, 1.6364),
		},
		{
			eye:    NewVector(0, 0, -1),
			normal: NewVector(0, 0, -1),
			light: PointLight{
				Position:  NewPoint(0, 0, 10),
				Intensity: NewColor(1, 1, 1),
			},
			result: NewColor(0.1, 0.1, 0.1),
		},
	}

	for _, test := range tests {
		got := material.Lighting(test.light, position, test.eye, test.normal)
		testColorEqual(t, got, test.result)
	}
}
