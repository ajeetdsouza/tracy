package engine

import "testing"

func TestShapeInterface(t *testing.T) {
	var shape Shape
	shape = &Sphere{}

	_ = shape
}
