package main

import (
	"fmt"
	"sync"
)

var wsg sync.WaitGroup

func main() {

	for i := 1; i <= 10; i++ {
		wsg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wsg.Done()
		}(i)
	}
	wsg.Wait()
}
