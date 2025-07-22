package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func multiplication(a, b int) {
	fmt.Printf("The multiplication of %d * %d is %d\n", a, b, a*b)
	wg.Done()
}

func main() {
	a, b := 2, 3
	wg.Add(2)

	go multiplication(a, b)

	go func() {
		fmt.Printf("The substraction of %d - %d = %d\n", b, a, b-a)
		wg.Done()
	}()

	wg.Wait()
}
