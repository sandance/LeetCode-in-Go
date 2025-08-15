package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	done := make(chan bool)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "result 1"
		time.Sleep(time.Second * 1)
		done <- true
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("received", msg1)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout one second after")
	case <-done:
		fmt.Println("done")
	}
}
