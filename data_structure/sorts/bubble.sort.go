package main

import "fmt"

func BubbleSort(arr []int, comp func(int, int) bool) {
	size := len(arr)

	for i := 0; i < (size - 1); i++ {
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func more(value1 int, value2 int) bool {
	return value1 > value2
}

func main() {
	data := []int{0, 15, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	BubbleSort(data, more)
	fmt.Println(data)
}
