package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	for i, num := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		result = search_pair(nums, -num, i+1, result)

	}

	// Replace this placeholder return statement with your code
	return result
}

func search_pair(nums []int, targetSum int, left int, result [][]int) [][]int {
	right := len(nums) - 1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == targetSum {
			result = append(result, []int{-targetSum, nums[left], nums[right]})
		}
		left++
		right--

		for left < right && nums[left] == nums[left-1] {
			left++
		}

		for left < right && nums[right] == nums[right+1] {
			right--
		}

	}
	return result
}

func main() {
	testCases := [][]int{
		[]int{-1, 0, 1, 2, -1, -4},
		[]int{-1, 0, 1, 2, -1, -4},
	}

	for _, testCase := range testCases {
		fmt.Printf("Input: %v, Output: %v\n", testCase, testCase)
		result := threeSum(testCase)
		fmt.Printf("Result: %v\n\n", result)
	}
}
