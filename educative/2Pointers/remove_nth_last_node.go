package main

import "fmt"

// Definition for a Linked List node

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthLastNode(head *ListNode, n int) *ListNode {

	// Replace this placeholder return statement with your code
	if head == nil {
		return nil
	}

	right := head
	left := head

	for i := 0; i < n; i++ {
		right = right.Next
	}

	if right == nil {
		return head.Next
	}

	for right.Next != nil {
		right = right.Next
		left = left.Next
	}

	left.Next = left.Next.Next

	return head
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	result := removeNthLastNode(head, 3)
	fmt.Println(result)

}
