package main

import "fmt"

func main() {
	goChannel := make(chan int)
	go func() {
		goChannel <- 1
	}()
	//close(goChannel)
	fmt.Println("Value: ", <-goChannel)
}
