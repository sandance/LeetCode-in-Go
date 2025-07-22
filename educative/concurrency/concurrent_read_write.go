package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func readMap(sharedMap map[int]int, mu *sync.Mutex) {
	for {
		mu.Lock()
		val := sharedMap[0]
		fmt.Println(val)
		mu.Unlock()
	}
}

func writeMap(sharedMap map[int]int, mu *sync.Mutex) {
	for {
		mu.Lock()
		sharedMap[0] = sharedMap[0] + 1
		mu.Unlock()
	}
}

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)
	sharedMap := make(map[int]int)
	sharedMap[0] = 0

	go readMap(sharedMap, &mu)
	go writeMap(sharedMap, &mu)

	wg.Wait()
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}
