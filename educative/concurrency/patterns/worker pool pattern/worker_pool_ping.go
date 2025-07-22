package main

import (
	"fmt"
	"log"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct {
	URL, workerIDMsg string
	Status           int
}

func pingWebsite(wId int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs {
		resp, err := http.Get(site.URL)
		if err != nil {
			log.Println(err.Error())
		}

		results <- Result{
			URL:         site.URL,
			workerIDMsg: fmt.Sprintf("\nThe worker id is %d, and status_code", wId),
			Status:      resp.StatusCode,
		}
	}
}

func main() {
	workersCount := 4
	jobs := make(chan Site, 3)      // buffered channel
	results := make(chan Result, 3) // buffered channel

	for w := 1; w <= workersCount; w++ {
		go pingWebsite(w, jobs, results)
	}

	urls := []string{
		"https://educative.io",
		"https://educative.io/learn",
		"https://educative.io/teach",
		"https://www.educative.io/explore/new",
		"https://www.educative.io/explore/picks",
		"https://www.educative.io/explore/early-access",
		"https://google.com",
	}

	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)

	for i := 1; i <= len(urls); i++ {
		result := <-results
		fmt.Println(result.URL, result.workerIDMsg)
	}
}
