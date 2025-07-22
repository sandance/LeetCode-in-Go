package main

import (
	"fmt"
	"strings"
)

func pow(digit int, power int) int {
	res := 1

	for i := 0; i < power; i++ {
		res = res * digit
	}
	return res
}

func sumOfSqaureDigits(n int) int {
	squareSum := 0
	for n > 0 {
		digit := n % 10
		squareSum += pow(digit, 2)
		n = n / 10
	}
	return squareSum
}

func isHappy(num int) bool {
	slowPointer := num
	fastPointer := sumOfSqaureDigits(slowPointer)

	for fastPointer != 1 && fastPointer != slowPointer {
		slowPointer = sumOfSqaureDigits(slowPointer)
		fastPointer = sumOfSqaureDigits(sumOfSqaureDigits(fastPointer))
	}
	if fastPointer == 1 {
		return true
	}
	return false
}

func main() {
	array := []int{1, 5, 19, 25, 7}
	for i, v := range array {
		fmt.Printf("%d.\tInput number: %d\n", i+1, v)
		fmt.Printf("\n\tIs it a happy number? %t\n", isHappy(v))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
