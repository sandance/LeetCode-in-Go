package main

import (
	"fmt"
	"sync"
)

func fibonacci(n int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		fmt.Println("Producer started")
		for i, j := 0, 1; i < n; i, j = j+i, i {
			output <- i
		}
		fmt.Println("Producer finished")
	}()
	return output
}

var wg sync.WaitGroup

func main() {
	// consumer
	for c := range fibonacci(10000) {
		fmt.Printf("Fibonacci number is : %d\n", c)
	}
}
