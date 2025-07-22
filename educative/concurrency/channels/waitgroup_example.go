package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	worker := func(wId int, wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Printf("Worker %d starting\n", wId)
	}
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

}
