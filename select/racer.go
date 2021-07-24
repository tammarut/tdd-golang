package racer

import (
	"fmt"
	"net/http"
	"time"
)

// 👉Racer(original)
// func Racer(urlA, urlB string) (winner string) {
// 	aDuration := measureResponseTime(urlA)
// 	bDuration := measureResponseTime(urlB)

// 	if aDuration < bDuration {
// 		return urlA
// 	}

// 	return urlB
// }

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	duration := time.Since(start)

// 	return duration
// }

// 👉Racer(Select)
var tenSecondTimeout = 10 * time.Second

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
		return "", fmt.Errorf("⏰Time out waiting for %s and %s", urlA, urlB)
	}
}

func ping(url string) chan struct{} {
	channel := make(chan struct{})
	go func(url string) {
		http.Get(url)
		close(channel)
	}(url)

	return channel
}
