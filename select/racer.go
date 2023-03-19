package main

import (
	"errors"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

func Racer(urlA, urlB string) (winner string, err error) {
	return ConfigurableRacer(urlA, urlB, tenSecondTimeout)
}

func ConfigurableRacer(urlA, urlB string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(timeout):
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

// Ditching in favor of time.After()
// func timeout() chan struct{} {
// 	ch := make(chan struct{})
// 	go func() {
// 		time.Sleep(10 * time.Second)
// 		close(ch)
// 	}()
// 	return ch
// }
