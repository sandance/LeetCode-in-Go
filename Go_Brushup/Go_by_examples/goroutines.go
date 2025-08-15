package main

import (
	"fmt"
	"sync"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {
	var wg sync.WaitGroup

	go f("hello from main")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// doing some work
			//fmt.Printf("working on values: %d\n", i)
			f(fmt.Sprintf("working on values %d", i))
		}(i)
	}

	//time.Sleep(time.Second)
	wg.Wait()
	fmt.Println("done")
}
