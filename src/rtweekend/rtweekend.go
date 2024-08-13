// Package rtweekend contains some utility functions
package rtweekend

import (
	"math"
	"math/rand"
)

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 100.0
}

func RandomFloat64() float64 {
	return rand.Float64()
}

func RandomFloat64InRange(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
