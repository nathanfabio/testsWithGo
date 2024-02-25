package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Nathan")
	expect := "Hello, Nathan"

	if result != expect {
		t.Errorf("result: %s, expect: %s", result, expect)
	}
}