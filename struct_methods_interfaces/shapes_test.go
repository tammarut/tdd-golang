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
	t.Run("rectangles", func(t *testing.T) {
		rectangle := rectangle{5.0, 2.0}
		got := rectangle.area()
		want := 10.0

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := circle{10}
		got := circle.area()
		want := 314.0

		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	})
}
