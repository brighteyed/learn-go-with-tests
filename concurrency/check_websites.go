package concurrency

// WebsiteChecker returns true if a given url is available
type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// CheckWebsites checks if urls in the given list are available
func CheckWebsites(w WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, w(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
