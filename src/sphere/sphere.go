// Package sphere contains the Sphere type and associated methods
package sphere

import (
	"math"

	"main/src/hittable"
	"main/src/point3"
	"main/src/ray"
	"main/src/vec3"
)

type Sphere struct {
	center point3.Point3
	radius float64
}

func New(p point3.Point3, radius float64) *Sphere {
	return &Sphere{p, radius}
}

func (s Sphere) Hit(r ray.Ray, rayMinT float64, rayMaxT float64, rec *hittable.HitRecord) bool {
	oc := s.center.AddVec3(r.Origin().ScaleFloat(-1.0))
	a := r.Direction().LengthSquared()
	h := vec3.Dot(r.Direction(), oc)
	c := oc.LengthSquared() - s.radius*s.radius

	discriminant := h*h - a*c
	if discriminant < 0 {
		return false
	}

	sqrtd := math.Sqrt(discriminant)
	root := (h - sqrtd) / a
	if root <= rayMinT || rayMaxT <= root {
		root = (h + sqrtd) / a
		if root <= rayMinT || rayMaxT <= root {
			return false
		}
	}

	rec.T = root
	rec.P = r.At(rec.T)
	// rec.Normal = rec.P.AddVec3(s.center.ScaleFloat(-1.0)).ScaleFloat(1 / s.radius)

	outwardNormal := (rec.P.AddVec3(s.center.ScaleFloat(-1.0))).ScaleFloat(1 / s.radius)
	rec.SetFaceNormal(r, outwardNormal)

	return true
}
