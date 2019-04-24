package engine

type Shape interface {
	GetTransform() Transform
	GetMaterial() Material
	Intersect(Ray) IntersectionList
	NormalAt(Tuple) Tuple
}
