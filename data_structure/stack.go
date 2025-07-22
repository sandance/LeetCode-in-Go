package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

/*
https://pkg.go.dev/github.com/golang-collections/collections/stack
Stack ADT Operations
Push(k) pushes a value k to the top of the stack.
Pop() removes an element from the top of the stack and return its value.
Top() returns the value of the element on top of the stack.
Size() returns the number of elements in the stack.
IsEmpty() tells us whether the stack is empty or not. If the stack is empty, it returns true. Otherwise,
it returns false.
*/

func main() {
	s := stack.New()
	s.Push(1)
	s.Push(2)
	fmt.Println("Top:", s.Peek()) // Should print 2
	fmt.Println("Pop:", s.Pop())  // Should print 2
	fmt.Println("IsEmpty:", s.Len() == 0)

	for s.Len() != 0 {
		val := s.Pop()
		fmt.Println(val, " ")
	}
}
