package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var ops uint64

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 100; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops: ", atomic.LoadUint64(&ops))
}
