package methodstruct

import "testing"

func TestPerimeter(t *testing.T) {
	got := perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := area(5.0, 2.0)
	want := 10.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}
