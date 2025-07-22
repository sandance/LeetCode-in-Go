package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type StackLinkedList struct {
	head *Node
	size int
}

// Size() function will return the size of the linked list.
func (s *StackLinkedList) Size() int {
	return s.Size() //Return 0 if stack is empty
}

/*
	IsEmpty() function will return true is size of the linked list is

equal to zero or false in all other cases.
*/
func (s *StackLinkedList) IsEmpty() bool {

	return s.size == 0 //Return true if stack is empty
}

/*
First, the Peek() function will check if the stack is empty.
If not, it will return the peek value of stack i.e., will return the
head value of the linked list.
*/
func (s *StackLinkedList) Peek() (int, bool) {
	//Implement your solution here
	if s.IsEmpty() {
		return 0, false
	}

	return (*s.head).value, true //Return 0,true if stack is empty
}

// Push() function  will add new value at the start of the linked list.
func (s *StackLinkedList) Push(value int) {
	//Implement your solution here
	s.head = &Node{
		value: value,
		next:  s.head,
	}
	s.size++
}

/*
In the pop() function, first it will check that the stack is not empty.
Then it will pop the value from the linked list and return it.
*/
func (s *StackLinkedList) Pop() (int, bool) {
	//Implement your solution here
	if s.IsEmpty() {
		return 0, false
	}

	value := s.head.value
	s.head = s.head.next
	s.size--
	return value, true
}

/* Print() function will print the elements of the linked list. */
func (s *StackLinkedList) Print() {
	//Implement your solution here
	temp := s.head
	fmt.Print("Value stored in stock are ::")
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

func main() {
	stack := new(StackLinkedList)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	val, _ := stack.Peek()
	fmt.Println("Peek of stack is", val)

	for stack.IsEmpty() == false {
		val, _ = stack.Pop()
		fmt.Println("Pop of stack is", val)
	}
}
