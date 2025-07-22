package main

import (
	"fmt"
	"time"
)

func worker(wId int, job <-chan int, results chan<- int) {
	for j := range job {
		fmt.Println("worker", wId, j)
		time.Sleep(time.Second)
		fmt.Println("finished worker", wId, j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		output := <-results
		fmt.Println(output)
	}
}
