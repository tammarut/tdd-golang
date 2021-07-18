package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wcCallback WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(url string) {
			resultChannel <- result{url, wcCallback(url)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		resultFromChannel := <-resultChannel
		results[resultFromChannel.string] = resultFromChannel.bool
	}

	return results
}
