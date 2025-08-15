package main

import (
	"fmt"
	"time"
)

func worker(wId int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("started Job", wId, j)
		time.Sleep(time.Second * 1)
		fmt.Println("finished Job", wId, j)
		results <- wId * j
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		<-results
	}

}
