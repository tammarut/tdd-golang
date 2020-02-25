package tutorial

import "testing"

func TestSquare(t *testing.T) {
	actual := Calculate(3)
	expected := 9

	if expected != actual {
		t.Errorf("expected %d but actual %d", expected, actual)
	}
}
