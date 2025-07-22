package main

import (
	"fmt"
	"strings"
)

func dutchFlagSort(arr []int) {
	low, high := 0, len(arr)-1

	current := 0

	for current <= high {
		if arr[current] == 0 {
			arr[current], arr[low] = arr[low], arr[current]
			current++
			low++
		} else if arr[current] == 1 {
			current++
		} else {
			arr[current], arr[high] = arr[high], arr[current]
			high--
		}
	}
}

func main() {
	inputs := [][]int{
		{0, 1, 0},
		{1, 1, 0, 2},
		{2, 1, 1, 0, 0},
		{2, 2, 2, 0, 1, 0},
		{2, 1, 1, 0, 1, 0, 2},
	}

	for i, input := range inputs {
		fmt.Printf("%d.\tcolors: %v\n", i+1, strings.Replace(fmt.Sprint(input), " ", ", ", -1))
		dutchFlagSort(input)
		fmt.Println(input)
	}
}
