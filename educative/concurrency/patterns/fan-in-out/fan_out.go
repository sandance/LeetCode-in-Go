package main

import (
	"fmt"
	"sync"
)

type Numbers struct {
	squared, wId int
}

var wg sync.WaitGroup

func generatePipeline(randomNum []int) <-chan int {
	out := make(chan int)
	defer close(out)

	for _, num := range randomNum {
		out <- num
	}
	return out
}

func squareNumber(in <-chan int, val int) <-chan Numbers {
	out := make(chan Numbers)
	defer close(out)

	go func(val int) {
		for n := range in {
			out <- Numbers{n * n, val}
		}

	}(val)
	return out
}

func displayData(squaredNums ...<-chan Numbers) {
	for _, num := range squaredNums {
		wg.Add(1)
		go func(value <-chan Numbers) {
			defer wg.Done()
			for val := range value {
				fmt.Printf("The squareed number is %d, and wId is %d\n", val.squared, val.wId)
			}
		}(num)
	}
}

func main() {
	randomNumbers := []int{13, 44, 56, 99, 9, 45, 67, 90, 78, 23}

	// generate the common channel with inputs
	inputChan := generatePipeline(randomNumbers)

	// Fan out to 3 go-routines
	c1 := squareNumber(inputChan, 1)
	c2 := squareNumber(inputChan, 2)
	c3 := squareNumber(inputChan, 3)

	displayData(c1, c2, c3)

	wg.Wait()
}
