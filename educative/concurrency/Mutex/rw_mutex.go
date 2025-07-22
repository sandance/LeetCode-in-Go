package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.RWMutex
	data := make(map[string]int)

	go func() {
		mu.Lock()
		data["key"] = 42
		mu.Unlock()
	}()

	//Reader
	mu.RLock()
	value := data["key"]
	mu.RUnlock()
	fmt.Println("Value: ", value)
}
