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

type StubStore struct {
	response string
}

type SpyStore struct {
	response  string
	cancelled bool
}

func (stub *StubStore) Fetch() string {
	time.Sleep(200 * time.Millisecond)
	return stub.response
}
func (stub *StubStore) Cancel() {
}

func (stub *SpyStore) Fetch() string {
	time.Sleep(200 * time.Millisecond)
	return stub.response
}

func (stub *SpyStore) Cancel() {
	stub.cancelled = true
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
	t.Run("Work normally", func(t *testing.T) {
		data := "Hello"
		stubStore := &StubStore{data}
		server := Server(stubStore)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`want %s, got %s`, response.Body.String(), data)
		}
	})

	t.Run("happy path that returns data from spyStore", func(t *testing.T) {
		data := "Hello happy"
		spyStore := &SpyStore{response: data}
		server := Server(spyStore)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`want %s, got %s`, response.Body.String(), data)
		}

		if spyStore.cancelled {
			t.Error("it should not have cacelled the store")
		}

	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "Hello"
		spyStore := &SpyStore{response: data}
		server := Server(spyStore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if !spyStore.cancelled {
			t.Error("spyStore was not told to cancel")
		}
	})
}
