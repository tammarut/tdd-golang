package tutorial

import "testing"

func TestSquare(t *testing.T) {
	shape := Square{side: 8}
	expected := 64.0

	assertCheckArea(t, expected, shape)
}

func TestTriangle(t *testing.T) {
	shape := Triangle{2.0, 5.0}
	expected := 5.0

	assertCheckArea(t, expected, shape)
}

func TestCircle(t *testing.T) {
	ball := Circle{7.0}
	expected := 147.0

	assertCheckArea(t, expected, ball)

}

func assertCheckArea(t *testing.T, expected float64, shape Shape) {
	t.Helper()
	actual := shape.CalculateArea()

	if expected != actual {
		t.Errorf("expected %.2f but actual %.2f", expected, actual)
	}
}
