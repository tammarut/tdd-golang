package syncness

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Increase()
		counter.Increase()
		counter.Increase()

		want := 3

		assertCounter(t, counter, want)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		counter := NewCounter()
		wantedCount := 1000

		var waitGroup sync.WaitGroup
		waitGroup.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Increase()
				waitGroup.Done()
			}()
		}
		waitGroup.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("want %d, but got %d", got.Value(), want)
	}
}
