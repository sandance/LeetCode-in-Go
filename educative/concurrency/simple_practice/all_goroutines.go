package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func multiplication(a, b int) {
	fmt.Printf("The muliplication of %d and %d is %d\n", a, b, a*b)
}
func main() {
	a := 2

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			multiplication(a, i)
			defer wg.Done()
		}(i)
		//go multiplication(a, i)
	}

	wg.Wait()

}
