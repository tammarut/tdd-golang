package methodstruct

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := rectangle{10.0, 10.0}
	got := perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

// func TestArea(t *testing.T) {
// 	checkArea := func(t *testing.T, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.area()
// 		if got != want {
// 			t.Errorf("got %g, want %g", got, want)
// 		}
// 	}

// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := rectangle{5.0, 2.0}
// 		checkArea(t, rectangle, 10.0)
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := circle{10}
// 		checkArea(t, circle, 314.0)
// 	})

// 	t.Run("triangle", func(t *testing.T) {
// 		triangle := triangle{5.0, 10.0}
// 		checkArea(t, triangle, 25.0)
// 	})

// }

func TestArea(t *testing.T) {
	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{rectangle{5.0, 2.0}, 10.0},
		{circle{10}, 314.0},
		{triangle{5.0, 10.0}, 25.0},
	}

	for _, tt := range areaTest {
		got := tt.shape.area()
		if got != tt.want {
			t.Errorf("got %g, want %g", got, tt.want)
		}
	}
}
