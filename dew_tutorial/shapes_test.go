package tutorial

import "testing"

func assertCheckArea(t *testing.T, expected float64, shape Shape) {
	t.Helper()
	actual := shape.CalculateArea()

	if expected != actual {
		t.Errorf("expected %.2f but actual %.2f", expected, actual)
	}
}

func TestCalculateArea(t *testing.T) {
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Triangle", shape: Triangle{2.0, 5.0}, hasArea: 5.0},
		{name: "Square", shape: Square{8.0}, hasArea: 64.0},
		{name: "Circle", shape: Circle{7.0}, hasArea: 147.0},
	}

	for _, v := range areaTest {
		t.Run(v.name, func(t *testing.T) {
			assertCheckArea(t, v.hasArea, v.shape)
		})
	}
}
