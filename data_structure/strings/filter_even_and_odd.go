package main

import "fmt"

func filter(slice []int) ([]int, []int) {
	even, odd := make([]int, 1), make([]int, 1)
	for _, value := range slice {
		if value%2 == 0 {
			even = append(even, value)
		} else {
			odd = append(odd, value)
		}

	}
	return even, odd
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("slice: ", slice)

	even, odd := filter(slice)
	fmt.Println("even: ", even)
	fmt.Println("odd: ", odd)
}
