package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
)

/*
https://pkg.go.dev/github.com/golang-collections/collections/queue
Queue ADT operations#
Add(k) adds an element k at the end of the queue.
Remove() removes the first element from the queue and returns its value.
Front() returns the first element of the queue.
Size() returns the number of elements in the queue.
IsEmpty() checks whether the queue is empty or not. If the queue is empty, it returns true. Otherwise, it returns false.
*/

func main() {
	s := queue.New()
	s.Enqueue(1)
	s.Enqueue(2)
	s.Enqueue(3)

	fmt.Println("IsEmpty:", s.Len() == 0)

	for s.Len() != 0 {
		val := s.Dequeue().(int)
		fmt.Println(val, " ")
	}
}
