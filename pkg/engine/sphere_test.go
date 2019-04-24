package engine

import (
	"math"
	"testing"
)

func TestSphereNormalAt(t *testing.T) {
	sphere := Sphere{
		Transform: NewTransform(),
	}

	tests := []struct {
		point, normal Tuple
	}{
		{NewPoint(1, 0, 0), NewVector(1, 0, 0)},
		{NewPoint(0, 1, 0), NewVector(0, 1, 0)},
		{NewPoint(0, 0, 1), NewVector(0, 0, 1)},
		{NewPoint(1/math.Sqrt(3), 1/math.Sqrt(3), 1/math.Sqrt(3)), NewVector(1/math.Sqrt(3), 1/math.Sqrt(3), 1/math.Sqrt(3))},
	}

	for _, test := range tests {
		testTupleEqual(t, sphere.NormalAt(test.point), test.normal)
	}
}
