package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func logMsg(msg string) {
	now := time.Now().Format("15:04:05")
	fmt.Printf("[%s] %s\n", now, msg)
}

func printResponses(resp []string) {
	output, _ := json.MarshalIndent(resp, "", "  ")
	logMsg(fmt.Sprintf("Got responses: %s", output))
}

// Simulates start_http_get in callback style
func startHTTPGet(url string, callback func(string)) {
	logMsg(fmt.Sprintf("HTTP start %q", url))

	go func() {
		resp, err := http.Get(url)
		if err != nil {
			callback(fmt.Sprintf("error: %v", err))
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		truncated := string(body)
		if len(truncated) > 20 {
			truncated = truncated[:20]
		}
		logMsg(fmt.Sprintf("HTTP finish %q -> %q", url, truncated))
		callback(truncated)
	}()
}

func startHTTPGet2(url string, callback func(string)) {
	logMsg(fmt.Sprintf("HTTP start %q", url))

	go func() {
		resp, err := http.Get(url)
		if err != nil {
			callback(fmt.Sprintf("error: %v", err))
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		truncated := string(body)
		if len(truncated) > 20 {
			truncated = truncated[:20]
		}
		logMsg(fmt.Sprintf("HTTP finish %q -> %q", url, truncated))
		callback(truncated)
	}()
}

func startHTTPGetSerial(urls []string, callback func([]string)) {
	responses := []string{}
	var runNext func(int)

	runNext = func(i int) {
		if i >= len(urls) {
			callback(responses)
			return
		}
		startHTTPGet(urls[i], func(body string) {
			responses = append(responses, body)
			runNext(i + 1)
		})
	}
	runNext(0)
}

func startHTTPGetParallel(urls []string, callback func([]string)) {
	responses := make([]string, len(urls))

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i, url := range urls {
		wg.Add(1)
		startHTTPGet(url, func(body string) {
			defer wg.Done()
			mu.Lock()
			responses[i] = body
			mu.Unlock()
		})
	}

	go func() {
		wg.Wait()
		callback(responses)
	}()

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go [serial|parallel]")
		return
	}
	mode := os.Args[1]
	urls := []string{
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/2",
		"https://httpbin.org/delay/1",
	}

	done := make(chan struct{})

	switch mode {
	case "serial":
		startHTTPGetSerial(urls, func(responses []string) {
			printResponses(responses)
			close(done)
		})
	case "parallel":
		startHTTPGetParallel(urls, func(responses []string) {
			printResponses(responses)
			close(done)
		})
	default:
		fmt.Printf("Invalid mode: %s\n", mode)
		return
	}

	logMsg("Started!")
	<-done
}
