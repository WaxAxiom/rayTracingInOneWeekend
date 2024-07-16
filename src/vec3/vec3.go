// Package vec3 contains the Vec3 type and supporting methods
package vec3

import (
	"math"
)

type Vec3 struct {
	e [3]float64
}

// Vector methods

func (v *Vec3) X() float64 {
	return v.e[0]
}

func (v *Vec3) Y() float64 {
	return v.e[1]
}

func (v *Vec3) Z() float64 {
	return v.e[2]
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v *Vec3) LengthSquared() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func New(e [3]float64) *Vec3 {
	return &Vec3{e}
}

func Sum(u Vec3, v Vec3) Vec3 {
	return Vec3{[3]float64{
		u.e[0] + v.e[0],
		u.e[1] + v.e[1],
		u.e[2] + v.e[2]}}
}

func ScaleFloat(v Vec3, t float64) Vec3 {
	return Vec3{[3]float64{t * v.e[0], t * v.e[1], t * v.e[2]}}
}

func ScaleVec3(u Vec3, v Vec3) Vec3 {
	return Vec3{[3]float64{u.e[0] * v.e[0], u.e[1] * v.e[1], u.e[2] * v.e[2]}}
}

func Dot(u Vec3, v Vec3) float64 {
	return u.e[0]*v.e[0] + u.e[1]*v.e[1] + u.e[2]*v.e[2]
}

func Cross(u Vec3, v Vec3) Vec3 {
	return Vec3{[3]float64{
		u.e[1]*v.e[2] - u.e[2]*v.e[1],
		u.e[2]*v.e[0] - u.e[0]*v.e[2],
		u.e[0]*v.e[1] - u.e[1]*v.e[0]}}
}

func UnitVector(v Vec3) Vec3 {
	return ScaleFloat(v, 1/v.Length())
}
