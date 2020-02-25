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
	actual := TriangleArea(2, 5)
	expected := 12

	if expected != actual {
		t.Errorf("expected %d but actual %d", expected, actual)
	}
}
