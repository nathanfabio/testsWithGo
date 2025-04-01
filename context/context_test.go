package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)


type SpyStore struct {
	response string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestHandler(t *testing.T) {
	data := "Hello World"
	svr := Server(&SpyStore{response: data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf("result '%s', expected '%s'", response.Body.String(), data)
	}

	t.Run("notify the store to cancell the work if request was cancelled", func(t *testing.T) {
		store := &SpyStore{response: data}
		svr := Server(store)


		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5 * time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if !store.cancelled {
			t.Errorf("store wasn't notified to cancell")
		}
	})

	t.Run("return data from store", func(t *testing.T) {
		store := SpyStore{response: data}
		svr := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("result '%s', expected '%s'", response.Body.String(), data)
		}

		if store.cancelled {
			t.Errorf("shouldn't have cancelled the store")
		}
	})
}
