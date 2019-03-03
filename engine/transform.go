package engine

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

type Transform struct {
	data mat.Matrix
}

func NewTransform() Transform {
	return Transform{
		mat.NewDense(4, 4, []float64{
			1, 0, 0, 0,
			0, 1, 0, 0,
			0, 0, 1, 0,
			0, 0, 0, 1,
		}),
	}
}

func (transform Transform) Transpose() Transform {
	return Transform{transform.data.T()}
}

func (transform Transform) Inverse() Transform {
	var result mat.Dense
	err := result.Inverse(transform.data)
	if err != nil {
		panic(fmt.Sprintf("Cannot compute inverse: %v", transform))
	}

	return Transform{&result}
}

func (transform Transform) Chain(other Transform) Transform {
	var result mat.Dense
	result.Mul(other.data, transform.data)
	return Transform{&result}
}

func (transform Transform) Translate(x, y, z float64) Transform {
	return transform.Chain(Transform{
		mat.NewDense(4, 4, []float64{
			1, 0, 0, x,
			0, 1, 0, y,
			0, 0, 1, z,
			0, 0, 0, 1,
		}),
	})
}

func (transform Transform) Scale(x, y, z float64) Transform {
	return transform.Chain(Transform{
		mat.NewDense(4, 4, []float64{
			x, 0, 0, 0,
			0, y, 0, 0,
			0, 0, z, 0,
			0, 0, 0, 1,
		}),
	})
}

func (transform Transform) RotateX(r float64) Transform {
	sinr, cosr := math.Sincos(r)
	return transform.Chain(Transform{
		mat.NewDense(4, 4, []float64{
			1, 0, 0, 0,
			0, cosr, -sinr, 0,
			0, sinr, cosr, 0,
			0, 0, 0, 1,
		}),
	})
}

func (transform Transform) RotateY(r float64) Transform {
	sinr, cosr := math.Sincos(r)
	return transform.Chain(Transform{
		mat.NewDense(4, 4, []float64{
			cosr, 0, sinr, 0,
			0, 1, 0, 0,
			-sinr, 0, cosr, 0,
			0, 0, 0, 1,
		}),
	})
}

func (transform Transform) RotateZ(r float64) Transform {
	sinr, cosr := math.Sincos(r)
	return transform.Chain(Transform{
		mat.NewDense(4, 4, []float64{
			cosr, -sinr, 0, 0,
			sinr, cosr, 0, 0,
			0, 0, 1, 0,
			0, 0, 0, 1,
		}),
	})
}

func (transform Transform) Shear(x_y, x_z, y_x, y_z, z_x, z_y float64) Transform {
	return transform.Chain(Transform{
		mat.NewDense(4, 4, []float64{
			1, x_y, x_z, 0,
			y_x, 1, y_z, 0,
			z_x, z_y, 1, 0,
			0, 0, 0, 1,
		}),
	})
}
