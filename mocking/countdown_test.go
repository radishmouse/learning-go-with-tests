package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{} // why do I need the address-of?
		spySleeper := &SpySleeper{}
		Countdown(buffer, spySleeper) // ah...because I'm passing a pointer...

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
		if spySleeper.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
		}
	})
	t.Run("sleep before every print", func(t *testing.T) {
		// I keep forgetting that I need the address of...
		// because I want a pointer to a new SpyCountdownOperations
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("got %v want %v", want, spySleepPrinter.Calls)
		}
	})
	t.Run("test configurable sleeper", func(t *testing.T) {
		sleepTime := 5 * time.Second

		spyTime := &SpyTime{} // need a pointer here, b/c spyTime.Sleep receives a pointer
		sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
		sleeper.Sleep()

		if spyTime.durationSlept != sleepTime {
			t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
		}
	})
}
