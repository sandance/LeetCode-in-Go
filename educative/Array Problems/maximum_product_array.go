package main

import "fmt"

// Brute Force
func maxProduct(nums []int) int {
	maxProduct := nums[0]

	n := len(nums)
	for i := 0; i < n; i++ {
		current := nums[i]
		maxProduct = max(maxProduct, current)
		for j := i + 1; j < n; j++ {
			current *= nums[j]
			maxProduct = max(maxProduct, current)
		}
	}
	return maxProduct

}

func main() {
	fmt.Println(maxProduct([]int{1, 2, 3, 4}))
	fmt.Println(maxProduct([]int{1, 2, -3, 4}))
}
