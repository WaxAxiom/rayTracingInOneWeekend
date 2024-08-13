// Package interval contains the Interval type and associated methods
package interval

type Interval struct {
	min float64
	max float64
}

func New(min float64, max float64) *Interval {
	return &Interval{min, max}
}

func (i Interval) Min() float64 {
	return i.min
}

func (i Interval) Max() float64 {
	return i.max
}

func (i Interval) Size() float64 {
	return i.max - i.min
}

func (i Interval) Contains(x float64) bool {
	return i.min <= x && x <= i.max
}

func (i Interval) Surrounds(x float64) bool {
	return i.min < x && x < i.max
}

func (i Interval) Clamp(x float64) float64 {
	if x < i.Min() {
		return i.Min()
	}
	if x > i.Max() {
		return i.Max()
	}
	return x
}
