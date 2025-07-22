package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func increment2(previous int) int {
	// pretend this is an expensive operation
	time.Sleep(time.Duration((rand.Intn(100) + 1)) * time.Millisecond)
	return previous + 1
}

func main() {
	noOfConsumers := 10

	data := make(chan int)
	go func() {
		defer close(data)
		for i := 0; i < 100; i++ {
			fmt.Printf("Processing %d\n", i)
			data <- increment2(i)
		}
	}()

	process := func(data <-chan int, wg *sync.WaitGroup, id int) {
		defer wg.Done()
		for inputData := range data {
			fmt.Printf("Consumer %d Value Printed by consumer %d\n", id, inputData)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < noOfConsumers; i++ {
		wg.Add(1)
		go process(data, &wg, i)
	}

	go func() {
		wg.Wait()
		//defer close(data)
	}()
	//wg.Wait()

	fmt.Println("All consumers finished.")
}
