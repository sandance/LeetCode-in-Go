package main

import "container/heap"

type Set struct {
	sum int
	n2  int
	n3  int
}

type minHeap []Set

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(Set))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].sum < h[j].sum
}

func newHeapFunc() *minHeap {
	min := &minHeap{}
	heap.Init(min)
	return min
}

func (h minHeap) Top() Set {
	return h[0]
}
