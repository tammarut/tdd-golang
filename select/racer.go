package racer

import (
	"net/http"
	"time"
)

func Racer(urlA, urlB string) (winner string) {
	aDuration := measureResponseTime(urlA)
	bDuration := measureResponseTime(urlB)

	if aDuration < bDuration {
		return urlA
	}

	return urlB
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	duration := time.Since(start)

	return duration
}
