package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl", b)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}
