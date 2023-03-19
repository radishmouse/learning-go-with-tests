package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("can test greet", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
