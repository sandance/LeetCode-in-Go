package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("Received messages", msg)
	default:
		fmt.Println("No messages received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("Sent messages", msg)
	default:
		fmt.Println("No messages sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("Received messages", msg)
	case sig := <-signals:
		fmt.Println("Got signal", sig)
	default:
		fmt.Println("No activity")
	}

}
