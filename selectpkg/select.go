package selectpkg

import (
	"net/http"
	"time"
)

func Runner(a, b string) (winner string) {
	timeA := responseTime(a)
	timeB := responseTime(b)

	if timeA < timeB {
		return a
	}

	return b
}

func responseTime(URL string) time.Duration {
	start := time.Now()
	http.Get(URL)
	return time.Since(start)
}
