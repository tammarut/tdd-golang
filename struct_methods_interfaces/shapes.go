package methodstruct

type Shape interface {
	area() float64
}
type rectangle struct {
	Width  float64
	Height float64
}
type circle struct {
	radius float64
}

type triangle struct {
	width  float64
	height float64
}

func perimeter(rec rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}

func (rec rectangle) area() float64 {
	return rec.Width * rec.Height
}

func (cir circle) area() float64 {
	return 3.14 * cir.radius * cir.radius
}

func (tri triangle) area() float64 {
	return 0.5 * tri.width * tri.height
}
