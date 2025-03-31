package selectpkg

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	t.Run("compares the speed of servers, returning the fastest one", func(t *testing.T) {
		serverSlow := delayedServer(20 * time.Millisecond)
		serverFast := delayedServer(0 * time.Millisecond)

		defer serverSlow.Close()
		defer serverFast.Close()

		URLSlow := serverSlow.URL
		URLFast := serverFast.URL

		expected := URLFast
		result, err := Runner(URLSlow, URLFast)
		if err != nil {
			t.Fatalf("didn't expect an error, but got one %v", err)
		}

		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})

	t.Run("return an error if server doesn't respond within 10s", func(t *testing.T) {
		server := delayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := Configurable(server.URL, server.URL, 20 * time.Millisecond)
		if err == nil {
			t.Error("was expecting an error, but I didn't get one")
		}
	})
}

func delayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}