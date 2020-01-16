package methodstruct

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := rectangle{10.0, 10.0}
	got := perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := rectangle{5.0, 2.0}
	got := area(rectangle)
	want := 10.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}
