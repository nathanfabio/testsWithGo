package sync

import (
	"sync"
	"testing"
)


func TestCounter(t *testing.T) {
	t.Run("increment three times", func(t *testing.T) {
		counter := Counter{}
		counter.Increment()
		counter.Increment()
		counter.Increment()

		checkCounter(t, counter, 3)
	})

	t.Run("runs concurrently safely", func(t *testing.T) {
		expectedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(expectedCount)

		for i := 0; i < expectedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Increment()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		checkCounter(t, counter, expectedCount)
	})
}

func checkCounter(t *testing.T, result Counter, expected int) {
	t.Helper()
	if result.Value() != expected {
		t.Errorf("result %d, expected %d", result.Value(), expected)
	}
}