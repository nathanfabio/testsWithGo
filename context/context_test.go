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
	t *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store was cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
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

// func (s *SpyStore) Cancel() {
// 	s.cancelled = true
// }

// func (s *SpyStore) assertWasCancelled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Errorf("store wasn't notified to cancell")
// 	}
// }

// func (s *SpyStore) assertWasNotCancelled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Errorf("store was notified to cancell")
// 	}
// }

func TestHandler(t *testing.T) {
	data := "Hello World"

	t.Run("notify the store to cancell the work if request was cancelled", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)


		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5 * time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)
	})

	t.Run("return data from store", func(t *testing.T) {
		store := SpyStore{response: data, t: t}
		svr := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("result '%s', expected '%s'", response.Body.String(), data)
		}
	})
}
