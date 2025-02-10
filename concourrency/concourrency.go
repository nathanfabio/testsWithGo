package concourrency


type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func WebsiteChecks(check WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	channelResult := make(chan result)

	for _, url := range urls {
		go func(u string) {
			channelResult <- result{u, check(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <- channelResult
		results[result.string] = result.bool
	}

	return results
}
