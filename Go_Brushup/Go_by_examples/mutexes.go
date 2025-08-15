package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	counters map[string]int
	mu       sync.Mutex
}

func (c *Counter) Inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
	fmt.Println("Inc:", name, c.counters[name])
}

func main() {
	c := Counter{
		counters: map[string]int{
			"one": 1,
			"two": 2,
		},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		defer wg.Done()
		for range n {
			c.Inc(name)
		}
	}

	wg.Add(3)
	go doIncrement("one", 1)
	go doIncrement("two", 2)
	go doIncrement("three", 3)

	wg.Wait()
	fmt.Println(c.counters)
}
