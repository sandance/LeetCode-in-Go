package main

type ListNode struct {
	val  int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func NewLinkedList(values []int) *LinkedList {
	ll := &LinkedList{}
	if len(values) > 0 {
		ll.createdLinkedList(values)
	}
	return ll
}

func (ll *LinkedList) createdLinkedList(values []int) {
	if len(values) == 0 {
		ll.head = nil
		return
	}

	ll.head = &ListNode{
		val: values[0],
	}

	current := ll.head
	for _, value := range values[1:] {
		current.next = &ListNode{
			val: value,
		}
		current = current.next
	}
}

func (ll *LinkedList) GetLength() int {
	length := 0
	current := ll.head

	for current != nil {
		length++
		current = current.next
	}
	return length
}
