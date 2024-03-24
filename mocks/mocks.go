package main

import (
	"bytes"
	"fmt"
)

func Countdown(output *bytes.Buffer) {
	fmt.Fprint(output, "3")
}