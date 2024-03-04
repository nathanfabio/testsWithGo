package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repetitions := Repeat("a")
	want := "aaaaa"

	if repetitions != want {
		t.Errorf("expected '%s' got '%s'", want, repetitions)
	}
}