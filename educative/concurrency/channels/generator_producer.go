package main

import (
	"fmt"
	"sync"
)

var wsg sync.WaitGroup

func main() {
	for val := range produce() {
		fmt.Println(val)
	}
}

func produce() chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		defer wsg.Wait()
		for i := 0; i < 10; i++ {
			wsg.Add(1)
			go func() {
				defer wsg.Done()
				out <- i
			}()
		}
	}()

	return out
}
