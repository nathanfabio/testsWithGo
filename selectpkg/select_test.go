package selectpkg

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	t.Run("return an error if server doesn't respond within 10s", func(t *testing.T) {
		serverA := delayedServer(11 * time.Second)
		serverB := delayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Runner(serverA.URL, serverB.URL)
		if err == nil {
			t.Errorf("error not found")
		}
	})


	serverSlow := delayedServer(20 * time.Millisecond)
	serverFast := delayedServer(0 * time.Millisecond)
	
	defer serverSlow.Close()
	defer serverFast.Close()

	URLSlow := serverSlow.URL
	URLFast := serverFast.URL

	expected := URLFast
	result, _ := Runner(URLSlow, URLFast)

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}

}

func delayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}