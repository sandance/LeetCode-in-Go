package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

type NumberRef struct {
	original, reverse int
}

func readFileRef(filename string) (<-chan int, error) {
	out := make(chan int)
	fp, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s\n", filename)
	}

	go func(file *os.File) {
		defer fp.Close()
		defer close(out)

		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			line = strings.TrimSuffix(line, "\n")
			line = strings.TrimSuffix(line, "\r")
			integer, _ := strconv.Atoi(line)
			out <- integer
			if err == io.EOF {
				break
			}
		}
	}(fp)

	return out, nil
}

// algorithm to reverse the number
func reverseNumberRef(n int) int {
	reverse := 0
	for n != 0 {
		remainder := n % 10
		reverse = reverse*10 + remainder
		n = n / 10
	}
	return reverse
}

func mergeRef(channels ...<-chan int) <-chan Number {
	result := make(chan Number)

	var mergeGroup sync.WaitGroup
	mergeGroup.Add(len(channels))

	send := func(c <-chan int) {
		for n := range c {
			result <- Number{
				original: n,
				reverse:  reverseNumber(n),
			}
		}
		mergeGroup.Done()
	}

	for _, c := range channels {
		go send(c)
	}

	go func() {
		mergeGroup.Wait()
		close(result)
	}()

	return result
}

func main() {
	channel1, err := readFile("/Users/rashed/ZOOMINFO/Code_Practice/LeetCode-in-Go/educative/concurrency/patterns/fan-in/file1.txt")
	if err != nil {
		fmt.Println("unable to open file1.txt")
	}

	channel2, err := readFile("/Users/rashed/ZOOMINFO/Code_Practice/LeetCode-in-Go/educative/concurrency/patterns/fan-in/file2.txt")
	if err != nil {
		fmt.Println("unable to open file2.txt")
	}

	channel3, err := readFile("/Users/rashed/ZOOMINFO/Code_Practice/LeetCode-in-Go/educative/concurrency/patterns/fan-in/file3.txt")
	if err != nil {
		fmt.Println("unable to open file3.txt")
	}

	results := merge(channel1, channel2, channel3)

	for result := range results {
		fmt.Printf("Original Number is %d, Reverse number is %d\n", result.original, result.reverse)
	}

}
