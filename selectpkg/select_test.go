package selectpkg

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	serverSlow := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	serverFast := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	

	URLSlow := serverSlow.URL
	URLFast := serverFast.URL

	expected := URLFast
	result := Runner(URLSlow, URLFast)

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}

	serverSlow.Close()
	serverFast.Close()
}