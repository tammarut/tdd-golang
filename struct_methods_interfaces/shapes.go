package methodstruct

type rectangle struct {
	Width  float64
	Height float64
}
type circle struct {
	radius float64
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
