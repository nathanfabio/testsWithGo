package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const lastWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Pause()
}

type SleeperConfigurable struct {
	duration time.Duration
	pause func(time.Duration)
}

func(s *SleeperConfigurable) Pause() {
	s.pause(s.duration)
}

func Countdown(output io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Pause()
		fmt.Fprintln(output, i)
	}

	sleeper.Pause()
	fmt.Fprint(output, lastWord)
}
func main() {
	sleeper := &SleeperConfigurable{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
