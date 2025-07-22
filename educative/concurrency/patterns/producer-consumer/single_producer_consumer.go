package main

import (
	"fmt"
	"math/rand"
	"time"
)

func increment(previous int) int {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	return previous + 1
}

func main() {
	data := make(chan int)

	// producer
	go func() {
		defer close(data)
		for i := 0; i < 10; i++ {
			data <- increment(i)
		}
	}()

	// consumer
	for i := range data {
		fmt.Printf("Value of i = %d\n", i)
	}
}
