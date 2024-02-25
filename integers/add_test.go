package integers

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(7, 7)
	want := 14

	if sum != want {
		t.Errorf("got %d want %d", sum, want)
	}
}