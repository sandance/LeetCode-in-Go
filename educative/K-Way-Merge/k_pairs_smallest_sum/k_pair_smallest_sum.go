package main

import (
	"fmt"
	"strings"
)

func kSmallestPairs(list1 []int, list2 []int, k int) [][]int {
	result := make([][]int, 0)
	minHeap := newHeapFunc()

	for i, num1 := range list1 {
		for j, num2 := range list2 {
			sum := num1 + num2
			if minHeap.Len() < k {
				minHeap.Push(Set{
					sum: sum,
					n2:  i,
					n3:  j,
				})
			} else {
				if minHeap.Top().sum < sum {
					break
				} else {
					_ = minHeap.Pop().(Set)
					minHeap.Push(Set{
						sum: sum,
						n2:  i,
						n3:  j,
					})
				}
			}

		}
	}

	for minHeap.Len() > 0 {
		poppedValue := minHeap.Pop().(Set)
		result = append(result, []int{poppedValue.n2, poppedValue.n3})
	}
	return result
}

// Driver code
func main() {
	list1 := [][]int{
		{2, 8, 9},
		{1, 2, 300},
		{1, 1, 2},
		{4, 6},
		{4, 7, 9},
		{1, 1, 2},
	}

	list2 := [][]int{
		{1, 3, 6},
		{1, 11, 20, 35, 300},
		{1, 2, 3},
		{2, 3},
		{4, 7, 9},
		{1},
	}
	k := []int{9, 30, 1, 2, 5, 4}

	for i, list := range list1 {
		fmt.Printf("%d.\t Input Pairs: %s,%s \n\tK = %d\n", i+1, strings.Replace(fmt.Sprint(list), " ", ", ", -1), strings.Replace(fmt.Sprint(list2[i]), " ", ", ", -1), k[i])
		fmt.Printf("\tPairs with the smallest sum are: %s\n", strings.Replace(fmt.Sprint(kSmallestPairs(list, list2[i], k[i])), " ", ", ", -1))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
