package main

import "math"

func minSubArrayLen(target int, nums []int) int {

	minSubArrayLen := math.MaxInt32
	currentWindowSum := 0
	windowStart := 0

	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		currentWindowSum += nums[windowEnd]

		for currentWindowSum >= target {
			currentWindowSum -= nums[windowStart]
			minSubArrayLen = min(minSubArrayLen, windowEnd-windowStart+1)
			windowStart++
		}
	}

	if minSubArrayLen == math.MaxInt32 {
		return 0 // no valid subarray found
	}

	// Replace this placeholder return statement with your code

	return minSubArrayLen
}
