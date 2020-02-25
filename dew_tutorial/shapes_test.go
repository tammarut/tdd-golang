package tutorial

import "testing"

func TestSquare(t *testing.T) {
	box := Square{side: 8}
	actual := box.CalculateArea()
	expected := 64

	if expected != actual {
		t.Errorf("expected %d but actual %d", expected, actual)
	}
}

func TestTriangle(t *testing.T) {
	triangle := Triangle{2.0, 5.0}
	actual := triangle.CalculateArea()
	expected := 5.0

	if expected != actual {
		t.Errorf("expected %f but actual %f", expected, actual)
	}
}

func TestCircle(t *testing.T) {
	ball := Circle{7.0}
	actual := ball.CalculateArea()
	expected := 147.0

	if expected != actual {
		t.Errorf("expected %.2f but actual %.2f", expected, actual)
	}
}
