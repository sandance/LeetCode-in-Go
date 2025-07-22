package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	noOfProducer := 10
	noOfConsumers := 10

	data := make(chan int64)

	var producerGroup sync.WaitGroup

	var ops int64
	for i := 0; i < noOfProducer; i++ {
		producerGroup.Add(1)
		go func() {
			defer producerGroup.Done()
			for c := 0; c < 100; c++ {
				data <- atomic.AddInt64(&ops, 1)
			}
		}()
	}

	go func() {
		producerGroup.Wait()
		defer close(data)
	}()

	var consumerGroup sync.WaitGroup
	for i := 0; i < noOfConsumers; i++ {
		consumerGroup.Add(1)
		go func() {
			defer consumerGroup.Done()
			for data := range data {
				fmt.Printf("Value Printed by consumer %d\n", data)
			}
		}()
	}
	consumerGroup.Wait()
}
