package concourrency



type WebsiteChecker func(string) bool


func WebsiteChecks(check WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = check(url)
	}

	return results
}
