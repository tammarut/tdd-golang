package contextness

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type Store interface {
	Fetch() string
	Cancel()
}

type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (stub *SpyStore) Fetch() string {
	time.Sleep(200 * time.Millisecond)
	return stub.response
}

func (stub *SpyStore) Cancel() {
	stub.cancelled = true
}

func (stub *SpyStore) assertWasCancelled() {
	stub.t.Helper()
	if !stub.cancelled {
		stub.t.Error("spyStore was not told to cancel")
	}
}

func (stub *SpyStore) assertWasOK() {
	stub.t.Helper()
	if stub.cancelled {
		stub.t.Error("spyStore should not have cacelled the store")
	}
}

func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case mydata := <-data:
			fmt.Fprint(writer, mydata)

		case <-ctx.Done():
			store.Cancel()
		}
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

		spyStore.assertWasOK()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "Hello"
		spyStore := &SpyStore{response: data, t: t}
		server := Server(spyStore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		spyStore.assertWasCancelled()
	})
}
