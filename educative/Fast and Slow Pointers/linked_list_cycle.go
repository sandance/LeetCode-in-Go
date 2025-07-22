package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func isCycle(head *Node) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	node1 := &Node{val: 1}
	node2 := &Node{val: 2}
	node3 := &Node{val: 3}
	node4 := &Node{val: 4}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node2 // cycle here

	if isCycle(node1) {
		fmt.Println("Cycle Detected")
	} else {
		fmt.Println("No Cycle Detected")
	}

}
