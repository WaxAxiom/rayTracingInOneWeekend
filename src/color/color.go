// Package color contains the Color type (an alias of Vec3) and supporting methods
package color

import (
	"fmt"
)

type Color struct {
	e [3]float64
}

func New(e [3]float64) Color {
	return Color{e}
}

func (c Color) R() float64 {
	return c.e[0]
}

func (c Color) G() float64 {
	return c.e[1]
}

func (c Color) B() float64 {
	return c.e[2]
}

func (v Color) AddFloat(t float64) Color {
	return Color{[3]float64{
		v.e[0] + t,
		v.e[1] + t,
		v.e[2] + t}}
}

func (v Color) AddColor(u Color) Color {
	return Color{[3]float64{
		v.e[0] + u.e[0],
		v.e[1] + u.e[1],
		v.e[2] + u.e[2]}}

}

func (v Color) ScaleFloat(t float64) Color {
	return Color{[3]float64{
		v.e[0] * t,
		v.e[1] * t,
		v.e[2] * t}}
}

func (v Color) ScaleColor(u Color) Color {
	return Color{[3]float64{
		v.e[0] * u.e[0],
		v.e[1] * u.e[1],
		v.e[2] * u.e[2]}}
}

func SumFloat(v Color, t float64) Color {
	return Color{[3]float64{
		t + v.e[0],
		t + v.e[1],
		t + v.e[2]}}
}

func SumColor(u Color, v Color) Color {
	return Color{[3]float64{
		u.e[0] + v.e[0],
		u.e[1] + v.e[1],
		u.e[2] + v.e[2]}}
}

func ScaleFloat(v Color, t float64) Color {
	return Color{[3]float64{t * v.e[0], t * v.e[1], t * v.e[2]}}
}

func ScaleColor(u Color, v Color) Color {
	return Color{[3]float64{u.e[0] * v.e[0], u.e[1] * v.e[1], u.e[2] * v.e[2]}}
}

func WriteColor(pixelColor Color) {
	ir := int(255.999 * pixelColor.R())
	ig := int(255.999 * pixelColor.G())
	ib := int(255.999 * pixelColor.B())

	fmt.Println(ir, ig, ib)
}
