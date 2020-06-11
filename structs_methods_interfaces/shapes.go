package main

type Shape interface {
	Area() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}
