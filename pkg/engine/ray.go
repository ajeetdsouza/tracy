package engine

type Ray struct {
	Origin, Direction Tuple
}

func (ray Ray) At(t float64) Tuple {
	return ray.Origin.Add(ray.Direction.Mul(t))
}

func (ray Ray) ApplyTransform(transform Transform) Ray {
	return Ray{
		ray.Origin.ApplyTransform(transform),
		ray.Direction.ApplyTransform(transform),
	}
}
