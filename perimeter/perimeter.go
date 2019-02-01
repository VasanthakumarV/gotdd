package perimeter

import "math"

type Shape interface {
	Area() float64
}

type Rectange struct {
	Width  float64
	Height float64
}

func (r *Rectange) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
