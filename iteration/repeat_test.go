package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := repeat("a", 5)
	expected := "aaaaa"

	if got != expected {
		t.Errorf("expected %s but got %s", expected, got)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := repeat("a", 4)
	fmt.Println(repeated)
	// Output: aaaa
}
