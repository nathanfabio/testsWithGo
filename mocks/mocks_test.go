package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountOperations{})

		result := buffer.String()
	expected := `3
2
1
Go!`

		if result != expected {
			t.Errorf("result '%s' expected '%s'", result, expected)
		}
	})

	t.Run("pause before every write", func(t *testing.T) {
		spySleepPrinting:= &SpyCountOperations{}
		Countdown(spySleepPrinting, spySleepPrinting)

		expected := []string{
			pause,
			write,
			pause,
			write,
			pause,
			write,
			pause,
			write,
		}

		if !reflect.DeepEqual(expected, spySleepPrinting.Calls) {
			t.Errorf("wanted calls %v got %v", expected, spySleepPrinting.Calls)
		}
	})
}

func TestSleeperConfigurable(t *testing.T) {
	pauseTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := &SleeperConfigurable{pauseTime, spyTime.Pause}
	sleeper.Pause()

	if spyTime.duration != pauseTime {
		t.Errorf("should have slept for %v but slept for %v", pauseTime, spyTime.duration)
	}
}

type SpyCountOperations struct {
	Calls []string
}

func (s *SpyCountOperations) Pause() {
	s.Calls = append(s.Calls, pause)
}

func (s *SpyCountOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const pause = "pause"

type SpyTime struct {
	duration time.Duration
}

func (t *SpyTime) Pause(duration time.Duration) {
	t.duration = duration
}