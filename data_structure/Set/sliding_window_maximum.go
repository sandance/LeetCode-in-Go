package main

import (
	"github.com/golang-collections/collections/queue"
)

func maxQueue(q queue.Queue) int {
	result := 0
	for q.Len() > 0 {
		if q.
	}
}

func findMaxSlidingWindow(nums []int, w int) []int {
	result := make([]int, 0)
	windowStart := 0

	queue := queue.New()

	for windowStart < len(nums) {
		rightChar := nums[windowStart]

		queue.Enqueue(rightChar)
		if queue.Len() == w {
			result = append(result, maxQueue(queue))
		}

	}

	// Replace this placeholder return statement with your code


	return result
}
