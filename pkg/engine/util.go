package engine

import (
	"testing"

	"github.com/ajeetdsouza/tracy/pkg/config"
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

func testInt64Equal(t *testing.T, got, exp int64) {
	if got != exp {
		t.Errorf("got %d, expected: %d", got, exp)
	}
}

func testFloat64Equal(t *testing.T, got, exp float64) {
	if !floats.EqualWithinAbsOrRel(got, exp, config.EPSILON, config.EPSILON) {
		t.Errorf("got %v, expected: %v", got, exp)
	}
}
