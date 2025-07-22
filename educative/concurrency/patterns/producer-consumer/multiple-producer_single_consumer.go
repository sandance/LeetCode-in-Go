package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	noOfProducer := 10

	data := make(chan int64)
	var wgz sync.WaitGroup

	var ops int64
	for i := 0; i < noOfProducer; i++ {
		wgz.Add(1)
		go func() {
			defer wgz.Done()
			for c := 0; c < 100; c++ {
				atomic.AddInt64(&ops, 1)
				data <- atomic.LoadInt64(&ops)
			}
		}()
	}

	go func() {
		wgz.Wait()
		close(data)
	}()

	// consumer
	for item := range data {
		fmt.Printf("Value of i = %d\n", item)
	}

}
