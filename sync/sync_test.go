package sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("increment three times", func(t *testing.T) {
		counter := Counter{}
			counter.Increment()
			counter.Increment()
			counter.Increment()

			if counter.Value() != 3 {
				t.Errorf("result %d, expected %d", counter.Value(), 3)
			}
	})	
}