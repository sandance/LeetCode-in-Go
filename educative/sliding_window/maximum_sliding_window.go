package main

import (
	"fmt"
	"github.com/gammazero/deque"
)

func findMaxSlidingWindow(nums []int, windowSize int) []int {
	var q deque.Deque[int]
	result := make([]int, 0)

	for _, num := range nums {

		q.PushBack(num)
		if q.Len() == windowSize {
			maxVal := findMaxValue(q)
			q.PopFront()
			result = append(result, maxVal)
		}
	}
	return result
}

func findMaxValue(q deque.Deque[int]) int {
	maxVal := q.At(0)
	for i := 1; i < q.Len(); i++ {
		if q.At(i) > maxVal {
			maxVal = q.At(i)
		}
	}
	return maxVal
}

func main() {
	fmt.Println(findMaxSlidingWindow([]int{9, 5, 3, 1, 6, 3}, 3))
}
