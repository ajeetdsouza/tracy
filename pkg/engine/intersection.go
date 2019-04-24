package engine

import "sort"

type Intersection struct {
	T     float64
	Shape Shape
}

type IntersectionList struct {
	data []Intersection
}

func NewIntersectionList(intersections ...Intersection) IntersectionList {
	sort.Slice(intersections, func(i, j int) bool {
		return intersections[i].T < intersections[j].T
	})
	return IntersectionList{intersections}
}

func (intersections IntersectionList) Hit() *Intersection {
	if len(intersections.data) > 0 {
		return &intersections.data[0]
	} else {
		return nil
	}
}
