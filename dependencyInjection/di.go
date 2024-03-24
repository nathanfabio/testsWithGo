package dependencyinjection

import (
	"fmt"
	"io"
)

func Greeting(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}