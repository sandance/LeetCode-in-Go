package main

import (
	"container/heap"
	"fmt"
)

/*
Heap ADT operations
Insert() adds a new element to the heap. Its time complexity is O(log(n))
Remove() extracts maximum value from a max-heap (or the minimum value from a min-heap). Its time complexity is O(log(n)).
Heapify() Converts a list of numbers in a list into a heap. Its time complexity is O(n)
*/

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	heap.Push(h, 0)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

}
