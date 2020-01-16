package methodstruct

type rectangle struct {
	Width  float64
	Height float64
}

func perimeter(rec rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}

func area(rec rectangle) float64 {
	return rec.Width * rec.Height
}
