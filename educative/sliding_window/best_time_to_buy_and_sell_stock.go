package main

import "fmt"

func maxProfit(prices []int) int {

	if len(prices) < 2 {
		return 0
	}
	buyIndex := 0
	maxProfit := 0

	for sellIndex := 1; sellIndex < len(prices); sellIndex++ {

		if prices[buyIndex] >= prices[sellIndex] {
			buyIndex = sellIndex
			continue
		}

		maxProfit = max(maxProfit, prices[sellIndex]-prices[buyIndex])
	}

	// Replace this placeholder return statement with your code

	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 5
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))    // 0
}
