package tutorial

import "testing"

func TestSquare(t *testing.T) {
	actual := Calculate(8)
	expected := 64

	if expected != actual {
		t.Errorf("expected %d but actual %d", expected, actual)
	}
}
