package main

import "fmt"

type fakeResult struct {
	body string
	urls []string
}
type fakeFetcher map[string]*fakeResult

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		body: "The Go Programming Language",
		urls: []string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		body: "Packages",
		urls: []string{
			"https://golang.org/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/cmd/": &fakeResult{
		body: "Commands",
		urls: []string{
			"https://golang.org/",
		},
	},
}
