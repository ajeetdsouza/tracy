package engine

type World struct {
	Shapes []Shape
}

func (world World) Intersect(ray Ray) IntersectionList {
	var intersections []Intersection
	for _, shape := range world.Shapes {
		intersections = append(intersections, shape.Intersect(ray).data...)
	}
	return NewIntersectionList(intersections...)
}
