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

// Sleeper of abstract
type Sleeper interface {
	Sleep()
}

// SpySleeper implement of Sleeper
type SpySleeper struct {
	Calls int
}

// Sleep method of SpySleeper
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// ConfigurableSleeper with a duration
type ConfigurableSleeper struct {
	duration time.Duration
}

// Sleep pause seconds with duration
func (c *ConfigurableSleeper) Sleep() {
	time.Sleep(c.duration)
}

// CountdownWithSleeper with a sleeper dependency
func CountdownWithSleeper(writer io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		s.Sleep()
		fmt.Fprint(writer, i)
	}

	s.Sleep()
	fmt.Fprint(writer, "Go!")
}
