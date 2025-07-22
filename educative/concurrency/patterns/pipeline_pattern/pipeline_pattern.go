package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Operation struct {
	id       int64
	multiply int64
	addition int64
}

func main() {
	c := displayData(prepareData(generateData()))

	for data := range c {
		log.Printf("Items: %v", data)
	}
}

// Stage 1: Read file and produce int64
func generateData() <-chan int64 {
	out := make(chan int64)
	const filePath = "/Users/rashed/ZOOMINFO/Code_Practice/LeetCode-in-Go/educative/concurrency/patterns/pipeline_pattern/integer.txt"

	go func() {
		defer close(out)

		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("Failed to open file: %v", err)
		}
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			line = strings.TrimSpace(line)
			if line != "" {
				if num, err := strconv.ParseInt(line, 10, 64); err == nil {
					out <- num
				}
			}
			if err == io.EOF {
				break
			}
		}
	}()

	return out
}

// Stage 2: Prepare data with multiple workers
func prepareData(in <-chan int64) <-chan Operation {
	out := make(chan Operation)
	var wg sync.WaitGroup

	worker := func() {
		defer wg.Done()
		for id := range in {
			out <- Operation{
				id:       id,
				multiply: id * 2,
				addition: id + 5,
			}
		}
	}

	numWorkers := 4
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// Stage 3: Display data (1 goroutine)
func displayData(in <-chan Operation) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for op := range in {
			out <- fmt.Sprintf("id: %d, multiply: %d, addition: %d", op.id, op.multiply, op.addition)
		}
	}()

	return out
}
