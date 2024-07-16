// Package ray contains the Ray type and supporting methods
package ray

import (
	"main/src/point3"
	"main/src/vec3"
)

type Ray struct {
	origin    point3.Point3
	direction vec3.Vec3
}

func (r Ray) Origin() point3.Point3 {
	return r.origin
}

func (r Ray) Direction() vec3.Vec3 {
	return r.direction
}

func (r Ray) At(t float64) point3.Point3 {
	return vec3.Sum(r.origin, vec3.ScaleFloat(r.direction, t))
}

func New(p point3.Point3, v vec3.Vec3) *Ray {
	return &Ray{p, v}
}
