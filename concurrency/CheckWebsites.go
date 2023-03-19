package main

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result) // a channel that holds result structs
	for _, url := range urls {
		go func(u string) {
			// Instead of putting the string,bool directly into the map...
			// results[u] = wc(u)
			// ...we'll push/write it to the channel
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// Then, we'll loop 0 to len(urls) and pull/read from the channel
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel

		// The values in a result struct are anonymouse
		// ...so, if it held two strings, how would you tell them apart?
		// yep. you can name them like so:
		// type result struct {
		// 	s: string
		// 	b: bool
		// }
		// allowing you to refer to them like so:
		// results[r.s] = r.b
		results[r.string] = r.bool
	}

	return results
}
