// Package point3 contains the Point3 type (an alias of Vec3) and supporting methods
package point

type Point3 struct {
	e [3]float64
}

func (p Point3) X() float64 {
	return p.e[0]
}

func (p Point3) Y() float64 {
	return p.e[1]
}

func (p Point3) Z() float64 {
	return p.e[2]
}

func New(e [3]float64) Point3 {
	return Point3{e}
}
