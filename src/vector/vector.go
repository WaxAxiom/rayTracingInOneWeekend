// Package vec3 contains the Vec3 type and supporting methods
package vector

import (
	"math"
)

type Vec3 struct {
	e [3]float64
}

func (v Vec3) New(e0 float64, e1 float64, e2 float64) {
	v.e[0] = e0
	v.e[1] = e1
	v.e[2] = e2
}

func (v Vec3) x() float64 {
	return v.e[0]
}

func (v Vec3) y() float64 {
	return v.e[1]
}

func (v Vec3) z() float64 {
	return v.e[2]
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) LengthSquared() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}
