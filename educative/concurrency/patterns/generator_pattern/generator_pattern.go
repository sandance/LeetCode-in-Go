package main

import (
	"fmt"
	"math/rand"
)

type result struct {
	value   int
	message string
}

// define struct

func main() {
	// write consumer function
	// you need to print the struct
	for returnVal := range randomGenerator(10) {
		fmt.Println(returnVal)
	}
}

// write a generator function which returns a channel
// call a goroutine to return random integer less than or equal to n
func randomGenerator(n int) <-chan result {
	output := make(chan result)

	go func() {
		defer close(output)
		fmt.Println("Hello, I am the producer")
		for i := 0; i < n; i++ {
			randNum := rand.Intn(n) + 1
			output <- result{
				value:   randNum,
				message: fmt.Sprintf("your random number is: %d", randNum),
			}
		}
	}()

	return output
}
