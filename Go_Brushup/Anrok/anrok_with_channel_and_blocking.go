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

func logMsg2(msg string) {
	now := time.Now().Format("15:04:05.000")
	fmt.Printf("[%s] %s\n", now, msg)
}

func startBlockingHTTPGet(url string) string {
	logMsg2(fmt.Sprintf("HTTP start %q", url))
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	truncated := string(body)
	if len(truncated) > 20 {
		truncated = truncated[:20]
	}
	logMsg2(fmt.Sprintf("HTTP finish %q -> %q", url, truncated))
	return truncated
}

// Serial execution using blocking calls
func startHTTPGetSerial2(urls []string) []string {
	var responses []string
	for _, url := range urls {
		res := startBlockingHTTPGet(url)
		responses = append(responses, res)
	}
	return responses
}

// Parallel execution using goroutines and channels
func startHTTPGetParallel2(urls []string) []string {
	type result struct {
		index int
		body  string
	}

	results := make(chan result, len(urls))
	//ch := make(chan result)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i, url := range urls {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			mu.Lock()
			body := startBlockingHTTPGet(url)
			results = append(results, result{i, body})
			//ch <- result{index: i, body: body}
		}(i, url)
	}

	responses := make([]string, len(urls))
	//for i := 0; i < len(urls); i++ {
	//	res := <-ch
	//	responses[res.index] = res.body
	//}
	for res := range ch {
		responses[res.index] = res.body
	}
	return responses
}

func printResponses2(responses []string) {
	out, _ := json.MarshalIndent(responses, "", "    ")
	logMsg2("Got responses: " + string(out))
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go [serial|parallel]")
		os.Exit(1)
	}
	mode := os.Args[1]

	urls := []string{
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/2",
		"https://httpbin.org/delay/1",
	}

	logMsg2("Started!")

	var responses []string
	switch mode {
	case "serial":
		responses = startHTTPGetSerial2(urls)
	case "parallel":
		responses = startHTTPGetParallel2(urls)
	default:
		fmt.Println("Expected 'serial' or 'parallel'")
		os.Exit(1)
	}

	printResponses2(responses)
}
