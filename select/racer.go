package main

import (
	"errors"
	"net/http"
	"time"
)

func Racer(urlA, urlB string) (winner string, err error) {

	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-timeout():
		return "", errors.New("ya took too long, now yo candy's gone, see")
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func timeout() chan struct{} {
	ch := make(chan struct{})
	go func() {
		time.Sleep(10 * time.Second)
		close(ch)
	}()
	return ch
}
