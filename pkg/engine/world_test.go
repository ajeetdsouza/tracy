package engine

import "testing"

func TestWorldIntersect(t *testing.T) {
	world := World{
		Shapes: []Shape{
			&Sphere{
				Material: Material{
					Color:     NewColor(0.8, 1.0, 0.6),
					Ambient:   0.0,
					Diffuse:   0.7,
					Specular:  0.2,
					Shininess: 0.0,
				},
				Transform: NewTransform(),
			},
			&Sphere{
				Material: Material{
					Color:     NewColor(0.0, 0.0, 0.0),
					Ambient:   0.0,
					Diffuse:   0.0,
					Specular:  0.0,
					Shininess: 0.0,
				},
				Transform: NewTransform().Scale(0.5, 0.5, 0.5),
			},
		},
	}

	ray := Ray{
		Origin:    NewPoint(0, 0, -5),
		Direction: NewVector(0, 0, 1),
	}

	intersections := world.Intersect(ray)
	testInt64Equal(t, int64(len(intersections.data)), 4)
	testFloat64Equal(t, intersections.data[0].T, 4.0)
	testFloat64Equal(t, intersections.data[1].T, 4.5)
	testFloat64Equal(t, intersections.data[2].T, 5.5)
	testFloat64Equal(t, intersections.data[3].T, 6.0)
}
