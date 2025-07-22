package main

import (
	"fmt"
	"time"
)

func countWithChannel(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}

// Experiencing channel deadlock.
func main() {
	c := make(chan string)
	//c <- "hello world" // Causes deadlock
	go countWithChannel("A", c)
	msg := <-c
	fmt.Println(msg)
}
