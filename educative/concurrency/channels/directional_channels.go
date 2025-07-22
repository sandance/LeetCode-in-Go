package main

import (
	"fmt"
	"sync"
)

var wag sync.WaitGroup

func printValue(goChannel <-chan int) {
	fmt.Println("Value inside the channel is ", <-goChannel)
	wag.Done()
}

func main() {
	goChannel := make(chan int)
	wag.Add(1)
	go printValue(goChannel)
	goChannel <- 2
	wag.Wait()
}

// send only channel
//sendOnly := make(chan<- int)
//receiveOnly := make(<-chan int)
