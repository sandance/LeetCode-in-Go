package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for c := range fibonnacci(1000) {
		fmt.Printf("Fibonacci number is %d \n", c)
	}
}

func fibonnacci(n int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()
	return out
}
