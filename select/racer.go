package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	startA := time.Now()
	http.Get(a)
	durationA := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	durationB := time.Since(startB)

	if durationA < durationB {
		return a
	}

	return b
}
