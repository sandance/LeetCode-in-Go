package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}

type Results struct {
	job              Job
	sumofdigits, wId int
}

func digits(number int) int {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit
		number /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func createWorkerPool(noOfWorkers int, jobs chan Job) chan Results {
	var wg sync.WaitGroup
	results := make(chan Results, 10)
	go func() {
		defer close(results)
		defer wg.Wait()
		for i := 0; i < noOfWorkers; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				for job := range jobs {
					output := Results{job, digits(job.randomno), i}
					results <- output
				}
			}(i)
		}
	}()

	//go func() {
	//	wg.Wait()
	//	close(results)
	//}()

	return results
}

func allocate(noOfJobs int) chan Job {
	jobs := make(chan Job, 10)
	go func() {
		defer close(jobs)
		for i := 0; i < noOfJobs; i++ {
			randomno := rand.Intn(999) + 1
			job := Job{i, randomno}
			jobs <- job
		}
	}()
	return jobs
}

func doWork(wId int, jobs <-chan Job, results chan<- Results, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		output := Results{job, digits(job.randomno), wId}
		results <- output
	}
}

func main() {
	jobs := make(chan Job, 100)
	results := make(chan Results, 100)

	noOfJobs := 100
	//jobGenerator := allocate(noOfJobs)
	noOfWorkers := 10

	var wg sync.WaitGroup

	for w := 1; w <= noOfWorkers; w++ {
		wg.Add(1)
		go doWork(w, jobs, results, &wg)
	}

	//for result := range createWorkerPool(noOfWorkers, jobGenerator) {
	//	fmt.Printf("Job id %d, input random no %d, sum of digits %d, worker id %d\n", result.job.id, result.job.randomno, result.sumofdigits, result.wId)
	//}

	for i := 0; i < noOfJobs; i++ {
		jobs <- Job{i, rand.Intn(999) + 1}
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	// Read results
	for output := range results {
		fmt.Printf("Worker %d processed JobID %d with input %d => sum of digits: %d\n",
			output.wId, output.job.id, output.job.randomno, output.sumofdigits)
	}

}
