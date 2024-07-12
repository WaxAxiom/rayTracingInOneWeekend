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

func WriteColor(pixelColor Color) {
	ir := int(255.999 * pixelColor.R())
	ib := int(255.999 * pixelColor.B())
	ig := int(255.999 * pixelColor.G())

	fmt.Println(ir, ib, ig)
}
