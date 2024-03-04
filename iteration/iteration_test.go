package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repetitions := Repeat("a", 5)
	want := "aaaaa" 

	if repetitions != want {
		t.Errorf("expected '%s' but got '%s'", want, repetitions)
	}
}


func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repetitions := Repeat("a", 10)
	fmt.Println(repetitions)
	//Output: aaaaaaaaaa
}