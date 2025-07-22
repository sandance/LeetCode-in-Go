package main

import "fmt"

func solveKnapsackRecursive(dp [][]int, profits []int, weights []int, capacity int, currentIndex int) int {
	if capacity == 0 || currentIndex >= len(profits) {
		return 0
	}

	if dp[currentIndex][capacity] != -1 {
		return dp[currentIndex][capacity]
	}

	profit1 := 0
	if weights[currentIndex] <= capacity {
		profit1 = profits[currentIndex] + solveKnapsackRecursive(dp, profits, weights, capacity-weights[currentIndex], currentIndex+1)
	}
	profit2 := solveKnapsackRecursive(dp, profits, weights, capacity, currentIndex+1)

	dp[currentIndex][capacity] = max(profit1, profit2)

	return dp[currentIndex][capacity]
}

func solveKnapsack(profits []int, weights []int, capacity int) int {
	dp := make([][]int, len(profits))
	for i := range dp {
		dp[i] = make([]int, capacity+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return solveKnapsackRecursive(dp, profits, weights, capacity, 0)
}

func main() {
	fmt.Println(solveKnapsack([]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 7))
	fmt.Println(solveKnapsack([]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 6))
}
