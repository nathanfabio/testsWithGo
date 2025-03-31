package selectpkg

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	serverSlow := delayedServer(20 * time.Millisecond)
	serverFast := delayedServer(0 * time.Millisecond)
	
	defer serverSlow.Close()
	defer serverFast.Close()

	URLSlow := serverSlow.URL
	URLFast := serverFast.URL

	expected := URLFast
	result := Runner(URLSlow, URLFast)

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