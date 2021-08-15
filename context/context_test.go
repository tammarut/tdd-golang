package contextness

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

type SpyStore struct {
	response string
	t        *testing.T
}

func (stub *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range stub.response {
			select {
			case <-ctx.Done():
				stub.t.Log("spy store got cancelled")
				return
			default:
				time.Sleep(100 * time.Millisecond)
				result += string(c)
			}
		}

		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func (stub *SpyStore) Cancel() {
	// stub.cancelled = true
}

type SpyResponseWriter struct {
	written bool
}

func (spyRes *SpyResponseWriter) Header() http.Header {
	spyRes.written = true
	return nil
}

func (spyRes *SpyResponseWriter) Write([]byte) (int, error) {
	spyRes.written = true
	return 0, errors.New("not implemented")
}

func (spyRes *SpyResponseWriter) WriteHeader(statusCode int) {
	spyRes.written = true
}

func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		data, err := store.Fetch(req.Context())
		if err != nil {
			return // TODO: log error
		}

		fmt.Fprint(writer, data)
	}
}

func TestServer(t *testing.T) {
	t.Run("happy path that returns data from spyStore", func(t *testing.T) {
		data := "Hello happy"
		spyStore := &SpyStore{response: data, t: t}
		server := Server(spyStore)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`want %s, got %s`, response.Body.String(), data)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "Hello cancel"
		spyStore := &SpyStore{response: data, t: t}
		server := Server(spyStore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		// Sadly httptest.NewRecorder doesn't have a way of figuring this out so we'll have to roll our own spy to test for this
		response := &SpyResponseWriter{}

		server.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})
}
