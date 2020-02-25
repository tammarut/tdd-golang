package tutorial

type Shape interface {
	CalculateArea() float64
}
type Square struct {
	side float64
}
type Triangle struct {
	height float64
	base   float64
}

type Circle struct {
	radius float64
}

func (s Square) CalculateArea() float64 {
	return s.side * s.side
}

func (t Triangle) CalculateArea() float64 {
	return 0.5 * t.height * t.base
}

func (c Circle) CalculateArea() float64 {
	return 22 / 7 * c.radius * c.radius
}
