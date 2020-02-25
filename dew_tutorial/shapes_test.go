package tutorial

import "testing"

func TestSquare(t *testing.T) {
	box := Square{side: 8.0}
	actual := box.CalculateArea()
	expected := 64.0

	assertCheckArea(t, expected, actual)
}

func TestTriangle(t *testing.T) {
	triangle := Triangle{2.0, 5.0}
	actual := triangle.CalculateArea()
	expected := 5.0

	assertCheckArea(t, expected, actual)
}

func TestCircle(t *testing.T) {
	ball := Circle{7.0}
	actual := ball.CalculateArea()
	expected := 147.0

	assertCheckArea(t, expected, actual)

}

func assertCheckArea(t *testing.T, expected, actual float64) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %.2f but actual %.2f", expected, actual)
	}
}
