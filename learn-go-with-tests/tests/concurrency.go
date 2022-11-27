package tests

// WebsiteChecker check website liveness
type WebsiteChecker func(string) bool

// CheckWebsites check multiple websites liveness status
func CheckWebsites(checker WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		results[url] = checker(url)
	}
	return results
}

type result struct {
	url    string
	status bool
}

// CheckWebsitesWithChannel check multiple websites liveness status with a channel
func CheckWebsitesWithChannel(checker WebsiteChecker, urls []string) map[string]bool {
	ch := make(chan result)
	results := make(map[string]bool)

	for _, url := range urls {
		go func(url string) {
			ch <- result{url: url, status: checker(url)}
		}(url)
	}

	for range urls {
		r := <-ch
		results[r.url] = r.status
	}

	return results
}
