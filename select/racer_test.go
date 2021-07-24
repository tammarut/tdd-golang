package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("success, return a fastest url", func(t *testing.T) {
		slowServer := newDelayedServer(100 * time.Millisecond)
		defer slowServer.Close()

		fastServer := newDelayedServer(0 * time.Millisecond)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)
		if err != nil {
			t.Fatalf("did not expect an error, but got %v", err)
		}

		if got != want {
			t.Errorf("got %q, but want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := newDelayedServer(25 * time.Millisecond)
		defer serverA.Close()

		_, err := ConfigurableRacer(serverA.URL, serverA.URL, 20*time.Millisecond)
		if err == nil {
			t.Error("Expected an error but didn't get one")
		}
	})
}

func newDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
