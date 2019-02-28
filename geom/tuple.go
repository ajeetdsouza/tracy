package geom

import (
	"fmt"

	"github.com/ajeetdsouza/tracy/config"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

const (
	wPoint  = 1.0
	wVector = 0.0
)

type Tuple struct {
	data mat.Vector
}

func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{mat.NewVecDense(4, []float64{x, y, z, w})}
}

func NewPoint(x, y, z float64) Tuple {
	return NewTuple(x, y, z, wPoint)
}

func NewVector(x, y, z float64) Tuple {
	return NewTuple(x, y, z, wVector)
}

func (tuple Tuple) X() float64 {
	return tuple.data.AtVec(0)
}

func (tuple Tuple) Y() float64 {
	return tuple.data.AtVec(1)
}

func (tuple Tuple) Z() float64 {
	return tuple.data.AtVec(2)
}

func (tuple Tuple) W() float64 {
	return tuple.data.AtVec(3)
}

func (tuple Tuple) String() string {
	return fmt.Sprintf("Tuple(%.1f,%.1f,%.1f,%.1f)", tuple.X(), tuple.Y(), tuple.Z(), tuple.W())
}

func (tuple Tuple) IsPoint() bool {
	return floats.EqualWithinAbsOrRel(tuple.W(), wPoint, config.EPSILON, config.EPSILON)
}

func (tuple Tuple) IsVector() bool {
	return floats.EqualWithinAbsOrRel(tuple.W(), wVector, config.EPSILON, config.EPSILON)
}

func (tuple Tuple) IsEqual(other Tuple) bool {
	return mat.EqualApprox(tuple.data, other.data, config.EPSILON)
}

func (tuple Tuple) Add(other Tuple) Tuple {
	var result mat.VecDense
	result.AddVec(tuple.data, other.data)
	return Tuple{&result}
}

func (tuple Tuple) Sub(other Tuple) Tuple {
	var result mat.VecDense
	result.SubVec(tuple.data, other.data)
	return Tuple{&result}
}

func (tuple Tuple) Mul(k float64) Tuple {
	var result mat.VecDense
	result.ScaleVec(k, tuple.data)
	return Tuple{&result}
}

func (tuple Tuple) Div(k float64) Tuple {
	var result mat.VecDense
	result.ScaleVec(1/k, tuple.data)
	return Tuple{&result}
}

func (tuple Tuple) Neg() Tuple {
	return tuple.Mul(-1)
}

func (tuple Tuple) Magnitude() float64 {
	return mat.Norm(tuple.data, 2)
}

func (tuple Tuple) Normalize() Tuple {
	r := tuple.Magnitude()
	return tuple.Div(r)
}

func (tuple Tuple) Dot(other Tuple) float64 {
	return mat.Dot(tuple.data, other.data)
}

func (tuple Tuple) Cross(other Tuple) Tuple {
	return NewVector(
		tuple.Y()*other.Z()-tuple.Z()*other.Y(),
		tuple.Z()*other.X()-tuple.X()*other.Z(),
		tuple.X()*other.Y()-tuple.Y()*other.X(),
	)
}
