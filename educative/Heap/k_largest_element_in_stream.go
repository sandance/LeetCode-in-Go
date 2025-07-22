package main

import (
	"container/heap"
	"fmt"
)

type KthLargest struct {
	topKHeap *MinHeap
	k        int
}

// Constructor to initialize heap and add values in it
func (kl *KthLargest) IKthLargestnit(k int) {
	fmt.Println("Initialiting the heap")
	kl.topKHeap = &MinHeap{}
	heap.Init(kl.topKHeap)
	kl.k = k

	fmt.Println("\tk =", kl.k)
	fmt.Println("\tHeap: []")

	for _, element := range nums {
		fmt.Println("Adding element", element, "to the heap.")
		kl.add(element)
	}
}

func (kl *KthLargest) add(element int) int {
	heap.Push(kl.topKHeap, element)
	if kl.topKHeap.Len() > kl.k {
		heap.Pop(kl.topKHeap)
	}
	return kl.topKHeap.Top()
}
