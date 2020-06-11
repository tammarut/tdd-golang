package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got: %.2f, want:%.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got: %.2f, want:%.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 2.0}
		want := 20.0

		checkArea(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.0

		checkArea(t, circle, want)
	})

	t.Run("triangle", func(t *testing.T) {
		triangle := Triangle{10.0, 5.0}
		want := 25.0

		checkArea(t, triangle, want)
	})
}
