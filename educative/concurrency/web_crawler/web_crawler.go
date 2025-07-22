package main

import (
	"fmt"
	"net/http"
	"time"
)

type webURL struct {
	URL string
}

func main() {

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	input := make(chan webURL)
	done := make(chan interface{})

	// producer
	webRequest := func(done <-chan interface{}, url string) {

		select {
		case <-done:
			return
		case input <- webURL{url}:
			time.Sleep(1 * time.Second)
		}

	}

	webResponse := func(done <-chan interface{}, urls <-chan webURL) {
		url := <-urls
		response, err := http.Get(url.URL)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Got response from %s\n", response.Request.URL)
	}

	urls := []string{
		"www.google.com",
		"www.facebook.com",
	}

	for _, url := range urls {
		go webRequest(done, url)
		go webResponse(done, input)
	}

}
