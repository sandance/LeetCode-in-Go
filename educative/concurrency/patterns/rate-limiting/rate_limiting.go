package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher, visited *sync.Map, wg *sync.WaitGroup, rateLimiter <-chan time.Time) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	if _, loaded := visited.LoadOrStore(url, true); loaded {
		return
	}

	<-rateLimiter

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, visited, wg, rateLimiter)
	}

}

func main() {
	var wg sync.WaitGroup
	visited := &sync.Map{}

	//create rate limiter 1 token per second
	rateLimiter := time.Tick(1 * time.Second)

	wg.Add(1)
	go Crawl("https://golang.org/", 3, fetcher, visited, &wg, rateLimiter)
	wg.Wait()
}
