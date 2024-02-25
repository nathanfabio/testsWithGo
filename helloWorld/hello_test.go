package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello()
	expect := "Hello, World!"

	if result != expect {
		t.Errorf("result: %s, expect: %s", result, expect)
	}
}