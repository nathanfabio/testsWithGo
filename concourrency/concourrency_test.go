package concourrency

import (
	"testing"
	"time"
)

func slowStubWebsiteVerification(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}


func BenchmarkWebsiteCheck(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "https://google.com"
	}

	for  i := 0; i < b.N; i++ {
		WebsiteChecks(slowStubWebsiteVerification, urls)
	}
}