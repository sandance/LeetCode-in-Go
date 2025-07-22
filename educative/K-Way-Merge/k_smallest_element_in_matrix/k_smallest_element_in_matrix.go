package main

import "container/heap"

func kthSmallestElement(matrix [][]int, k int) int {

	// Replace this placeholder return statement with your code
	minHeap := newMinHeap()

	for _, list := range matrix {
		heap.Push(minHeap, Set{
			list[0],
			0,
			list,
		})
	}

	numberCount := 0

	for !minHeap.IsEmpty() {
		poppedValue := heap.Pop(minHeap).(Set)
		numberCount++

		if numberCount == k {
			return poppedValue.num
		}

		if len(poppedValue.list) > poppedValue.index+1 {
			heap.Push(minHeap, Set{
				num:   poppedValue.list[poppedValue.index+1],
				index: poppedValue.index + 1,
				list:  poppedValue.list,
			})
		}
	}

	return -1
}
