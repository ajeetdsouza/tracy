package engine

import "math"

type Material struct {
	Color                                 Color
	Ambient, Diffuse, Specular, Shininess float64
}

func (material Material) Lighting(light PointLight, point, eye, normal Tuple) Color {
	// combine the surface color with the light's color/intensity
	effectiveColor := material.Color.MulColor(light.Intensity)

	// find the direction to the light source
	lightVector := light.Position.Sub(point).Normalize()

	// compute the ambient contribution
	ambient := effectiveColor.MulScalar(material.Ambient)

	// lightDotNormal represents the cosine of the angle between the
	// light vector and the normal vector. A negative number means the
	// light is on the other side of the surface.
	lightDotNormal := lightVector.Dot(normal)

	var diffuse, specular Color
	if lightDotNormal < 0 {
		diffuse = NewColor(0, 0, 0)
		specular = NewColor(0, 0, 0)
	} else {
		// compute the diffuse contribution
		diffuse = effectiveColor.MulScalar(material.Diffuse * lightDotNormal)

		// reflectDotEye represents the cosine of the angle between the
		// reflection vector and the eye vector. A negative number means the
		// light reflects away from the eye.
		reflectVector := lightVector.Neg().Reflect(normal)
		reflectDotEye := reflectVector.Dot(eye)

		if reflectDotEye <= 0 {
			specular = NewColor(0, 0, 0)
		} else {
			// compute the specular contribution
			factor := math.Pow(reflectDotEye, material.Shininess)
			specular = light.Intensity.MulScalar(material.Specular * factor)
		}
	}

	// Add the three contributions together to get the final shading
	return ambient.Add(diffuse).Add(specular)
}
