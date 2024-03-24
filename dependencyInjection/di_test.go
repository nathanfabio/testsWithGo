package main

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {
	buffer := bytes.Buffer{}
	Greeting(&buffer, "Nathan")


	result:= buffer.String()
	expected:= "Hello, Nathan"

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}