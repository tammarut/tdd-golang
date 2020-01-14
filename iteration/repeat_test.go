package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := repeat("a")
	expected := "aaaaa"

	if got != expected {
		t.Errorf("expected %s but got %s", expected, got)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("a")
	}
}
