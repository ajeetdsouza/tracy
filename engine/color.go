package engine

import (
	"github.com/ajeetdsouza/tracy/config"
	"gonum.org/v1/gonum/mat"
)

type Color struct {
	data mat.Vector
}

func NewColor(r, g, b float64) Color {
	return Color{mat.NewVecDense(3, []float64{r, g, b})}
}

func (color Color) R() float64 {
	return color.data.AtVec(0)
}

func (color Color) G() float64 {
	return color.data.AtVec(1)
}

func (color Color) B() float64 {
	return color.data.AtVec(2)
}

func (color Color) Equal(other Color) bool {
	return mat.EqualApprox(color.data, other.data, config.EPSILON)
}

func (color Color) Add(other Color) Color {
	var result mat.VecDense
	result.AddVec(color.data, other.data)
	return Color{&result}
}

func (color Color) Sub(other Color) Color {
	var result mat.VecDense
	result.SubVec(color.data, other.data)
	return Color{&result}
}

func (color Color) MulColor(other Color) Color {
	var result mat.VecDense
	result.MulElemVec(color.data, other.data)
	return Color{&result}
}

func (color Color) MulScalar(k float64) Color {
	var result mat.VecDense
	result.ScaleVec(k, color.data)
	return Color{&result}
}
