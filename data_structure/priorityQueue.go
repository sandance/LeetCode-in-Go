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

type Item struct {
	value    string
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	items := map[string]int{
		"banana": 3,
		"apple":  2,
		"pear":   4,
	}

	pq := make(PriorityQueue, len(items))

	i := 0
	for value, priority := range items {
		pq[i] = &Item{value, priority, i}
		i++
	}

	heap.Init(&pq)

	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Println(item.value, item.priority)
	}
}
