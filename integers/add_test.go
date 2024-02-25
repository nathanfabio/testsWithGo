package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(7, 7)
	want := 14

	if sum != want {
		t.Errorf("got %d want %d", sum, want)
	}
}

func ExampleAdd() {
	sum := Add(6, 10)
	fmt.Println(sum)
	// Output: 16
}