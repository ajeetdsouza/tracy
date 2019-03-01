package geom

import (
	"math"
	"testing"

	"github.com/ajeetdsouza/tracy/config"
	"gonum.org/v1/gonum/floats"
)

func testTupleEqual(t *testing.T, got, exp Tuple) {
	if !got.IsEqual(exp) {
		t.Errorf("got %v, expected: %v", got, exp)
	}
}

func testFloatEqual(t *testing.T, got, exp float64) {
	if !floats.EqualWithinAbsOrRel(got, exp, config.EPSILON, config.EPSILON) {
		t.Errorf("got %v, expected: %v", got, exp)
	}
}

func TestPoint(t *testing.T) {
	tuple := NewTuple(4.3, -4.2, 3.1, 1.0)
	if !tuple.IsPoint() || tuple.IsVector() {
		t.Errorf("got vector, expected point")
	}
}

func TestVector(t *testing.T) {
	tuple := NewTuple(4.3, -4.2, 3.1, 0.0)
	if !tuple.IsVector() || tuple.IsPoint() {
		t.Error("got point, expected vector")
	}
}

func TestNewPoint(t *testing.T) {
	got := NewPoint(4, -4, 3)
	exp := NewTuple(4, -4, 3, 1)
	testTupleEqual(t, got, exp)
}

func TestNewVector(t *testing.T) {
	got := NewVector(4, -4, 3)
	exp := NewTuple(4, -4, 3, 0)
	testTupleEqual(t, got, exp)
}

func TestTupleAdd(t *testing.T) {
	a := NewPoint(3, -2, 5)
	b := NewVector(-2, 3, 1)

	got := a.Add(b)
	exp := NewTuple(1, 1, 6, 1)

	testTupleEqual(t, got, exp)
}

func TestTupleSub(t *testing.T) {
	a := NewPoint(3, 2, 1)
	b := NewVector(5, 6, 7)

	got := a.Sub(b)
	exp := NewPoint(-2, -4, -6)

	testTupleEqual(t, got, exp)
}

func TestTupleNeg(t *testing.T) {
	a := NewTuple(1, -2, 3, -4)

	got := a.Neg()
	exp := NewTuple(-1, 2, -3, 4)

	testTupleEqual(t, got, exp)
}

func TestTupleMul(t *testing.T) {
	a := NewTuple(1, -2, 3, -4)

	got := a.Mul(3.5)
	exp := NewTuple(3.5, -7, 10.5, -14)

	testTupleEqual(t, got, exp)
}

func TestTupleDiv(t *testing.T) {
	a := NewTuple(1, -2, 3, -4)

	got := a.Div(2)
	exp := NewTuple(0.5, -1, 1.5, -2)

	testTupleEqual(t, got, exp)
}

func TestTupleMagnitude(t *testing.T) {
	tests := []struct {
		tuple     Tuple
		magnitude float64
	}{
		{NewVector(1, 0, 0), 1.0},
		{NewVector(0, 1, 0), 1.0},
		{NewVector(0, 0, 1), 1.0},
		{NewVector(1, 2, 3), math.Sqrt(14)},
		{NewVector(-1, -2, -3), math.Sqrt(14)},
	}

	for _, test := range tests {
		testFloatEqual(t, test.tuple.Magnitude(), test.magnitude)
	}
}

func TestTupleNormalize(t *testing.T) {
	tests := []struct {
		tuple      Tuple
		normalized Tuple
	}{
		{NewVector(4, 0, 0), NewVector(1, 0, 0)},
		{NewVector(1, 2, 3), NewVector(0.26726, 0.53452, 0.80178)},
	}

	for _, test := range tests {
		testTupleEqual(t, test.tuple.Normalize(), test.normalized)
	}
}

func TestTupleNormalizedMagnitude(t *testing.T) {
	tuple := NewVector(1, 2, 3)
	got := tuple.Normalize().Magnitude()
	exp := 1.0
	testFloatEqual(t, got, exp)
}

func TestTupleDot(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)

	got := a.Dot(b)
	exp := 20.0
	testFloatEqual(t, got, exp)
}

func TestTupleCross(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)

	got := a.Cross(b)
	exp := NewVector(-1, 2, -1)
	testTupleEqual(t, got, exp)

	got = b.Cross(a)
	exp = NewVector(1, -2, 1)
	testTupleEqual(t, got, exp)
}
