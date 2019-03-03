package engine

import "math"

type Sphere struct {
	Material  Material
	Transform Transform
}

func (sphere *Sphere) Intersect(ray Ray) IntersectionList {
	ray = ray.ApplyTransform(sphere.GetTransform().Inverse())
	sphereToRay := ray.Origin.Sub(NewPoint(0, 0, 0))

	a := ray.Direction.Dot(ray.Direction)
	b := 2 * ray.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return NewIntersectionList()
	}

	i1 := Intersection{
		T:     (-b - math.Sqrt(discriminant)) / (2 * a),
		Shape: sphere,
	}
	i2 := Intersection{
		T:     (-b + math.Sqrt(discriminant)) / (2 * a),
		Shape: sphere,
	}

	return NewIntersectionList(i1, i2)
}

func (sphere *Sphere) NormalAt(worldPoint Tuple) Tuple {
	transform := sphere.GetTransform().Inverse()
	objectPoint := worldPoint.ApplyTransform(transform)

	objectNormal := objectPoint.Sub(NewPoint(0, 0, 0))

	transform = transform.Transpose()
	worldNormal := objectNormal.ApplyTransform(transform)

	return NewVector(worldNormal.X(), worldNormal.Y(), worldNormal.Z()).Normalize()
}

func (sphere *Sphere) GetMaterial() Material {
	return sphere.Material
}

func (sphere *Sphere) GetTransform() Transform {
	return sphere.Transform
}
