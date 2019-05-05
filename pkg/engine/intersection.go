package engine

import "sort"

type Intersection struct {
	T     float64
	Shape Shape
}

type IntersectionList struct {
	data []Intersection
}

// NewIntersectionList creates a sorted list of intersections
func NewIntersectionList(intersections ...Intersection) IntersectionList {
	sort.Slice(intersections, func(i, j int) bool {
		return intersections[i].T < intersections[j].T
	})
	return IntersectionList{intersections}
}

// Hit returns the nearest intersection (or nil, in case there is no intersection)
func (intersections IntersectionList) Hit() *Intersection {
	if len(intersections.data) > 0 {
		return &intersections.data[0]
	}
	return nil
}
