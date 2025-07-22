package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	//Note that the total execution time is only ~2 seconds since both the 1 and 2 second Sleeps execute concurrently.
	for range 2 {
		select {
		case msg := <-c1:
			fmt.Println("Received:", msg)
		case msg2 := <-c2:
			fmt.Println("Received:", msg2)
		}

	}
}
