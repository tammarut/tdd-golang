package arrayslice

import "testing"

func TestSum(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d by given %v", got, want, numbers)
		}
}
