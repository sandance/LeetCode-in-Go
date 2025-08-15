package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiters := time.Tick(200 * time.Millisecond)
	for req := range requests {
		<-limiters
		fmt.Println("request", req, time.Now())
	}

	brustyLimiter := make(chan time.Time, 3)

	for range 3 {
		brustyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			brustyLimiter <- t
		}
	}()

	brustyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		brustyRequests <- i
	}
	close(brustyRequests)

	for req := range brustyRequests {
		<-brustyLimiter
		fmt.Println("request", req, time.Now())
	}
}
