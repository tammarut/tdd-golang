package racer

import (
	"net/http"
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
func Racer(urlA, urlB string) (winner string) {
	select {
	case <-ping(urlA):
		return urlA
	case <-ping(urlB):
		return urlB
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
