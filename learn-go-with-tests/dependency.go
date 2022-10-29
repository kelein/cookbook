package tests

import (
	"fmt"
	"io"
	"time"
)

// Greet welcome to everyone
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// Countdown calculates time elapsed in seconds
func Countdown(writer io.Writer, seconds int) {
	for i := seconds; i > 0; i-- {
		time.Sleep(time.Second)
		fmt.Fprintf(writer, "%d\n", i)
	}
	fmt.Fprint(writer, "Go!")
}
