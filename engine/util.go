package engine

import (
	"testing"

	"github.com/ajeetdsouza/tracy/config"
	"gonum.org/v1/gonum/floats"
)

func testTupleEqual(t *testing.T, got, exp Tuple) {
	if !got.Equal(exp) {
		t.Errorf("got %v, expected: %v", got, exp)
	}
}

func testColorEqual(t *testing.T, got, exp Color) {
	if !got.Equal(exp) {
		t.Errorf("got %v, expected: %v", got, exp)
	}
}

func testFloatEqual(t *testing.T, got, exp float64) {
	if !floats.EqualWithinAbsOrRel(got, exp, config.EPSILON, config.EPSILON) {
		t.Errorf("got %v, expected: %v", got, exp)
	}
}
