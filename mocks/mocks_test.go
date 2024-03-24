package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	result:= buffer.String()
	expected:= "3"

	if result != expected {
		t.Errorf("result '%s', ecpected '%s'", result, expected)
	}
}