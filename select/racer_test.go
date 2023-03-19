package main

import "testing"

func TestRacer(t *testing.T) {
	t.Run("will it blend", func(t *testing.T) {
		slowURL := "http://www.facebook.com"
		fastURL := "http://quii.dev"
		want := fastURL
		got := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}
