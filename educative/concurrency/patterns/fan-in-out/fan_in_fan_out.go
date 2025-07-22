package main

import (
	"fmt"
	"runtime"
	"sync"
)

func worker(id int, urls <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range urls {
		results <- url
	}
}

var wkg sync.WaitGroup

func main() {
	urlsList := []string{
		"https://google.com",
		"https://golang.org",
		"https://github.com",
		"https://example.com",
	}

	urls := make(chan string, len(urlsList))
	results := make(chan string, len(urlsList))

	noOfWorkers := runtime.GOMAXPROCS(4)

	for i := 0; i < noOfWorkers; i++ {
		wkg.Add(1)
		go worker(i, urls, results, &wkg)
	}

	// producer
	for _, url := range urlsList {
		urls <- url
	}

	defer close(urls)

	wkg.Wait()
	close(results)

	// consumer
	for r := range results {
		fmt.Println(r)
	}
}
