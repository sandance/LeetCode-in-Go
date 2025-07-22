package main

import (
	"container/heap"
	"fmt"
)

//type ListNode struct {
//	value int
//	next  *ListNode
//}

func mergeKSortedLists(lists []*ListNode) *ListNode {
	minHeap := newMinHeap()

	// Replace this placeholder return statement with your code
	for _, list := range lists {
		if list != nil {
			heap.Push(minHeap, list)
		}
	}

	var head, tail *ListNode

	for !minHeap.IsEmpty() {
		node := heap.Pop(minHeap).(*ListNode)
		if head == nil {
			head = node
			tail = node
		} else {
			tail.next = node
			tail = tail.next
		}

		if node.next != nil {
			heap.Push(minHeap, node.next)
		}
	}
	return head
}

// printList prints a linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.value)
		head = head.next
	}
	fmt.Println("nil")
}

// buildList builds a linked list from a slice of ints
func buildList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{value: values[0]}
	current := head
	for _, val := range values[1:] {
		current.next = &ListNode{value: val}
		current = current.next
	}
	return head
}

// main runs a test
func main() {
	list1 := buildList([]int{1, 4, 7})
	list2 := buildList([]int{2, 5, 8})
	list3 := buildList([]int{3, 6, 9})

	lists := []*ListNode{list1, list2, list3}

	merged := mergeKSortedLists(lists)
	fmt.Println("Merged Sorted List:")
	printList(merged)
}
