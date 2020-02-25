package tutorial

import "testing"

func TestSquare(t *testing.T) {
	actual := Calculate(8)
	expected := 64

	if expected != actual {
		t.Errorf("expected %d but actual %d", expected, actual)
	}
}

func TestTriangle(t *testing.T) {
	actual := TriangleArea(2.0, 5.0)
	expected := 5.0

	if expected != actual {
		t.Errorf("expected %f but actual %f", expected, actual)
	}
}

func TestCircle(t *testing.T) {
	actual := CircleArea(7.0)
	expected := 154.0

	if expected != actual {
		t.Errorf("expected %f but actual %f", expected, actual)
	}
}
