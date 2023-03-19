package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("can do a thing", func(t *testing.T) {
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
}
