package main

import "fmt"

func calculateFibonacciRecurseive(memorize []int, n int) int {
	if n < 2 {
		return n
	}

	if memorize[n] != -1 {
		return memorize[n]
	}

	memorize[n] = calculateFibonacciRecurseive(memorize, n-1) + calculateFibonacciRecurseive(memorize, n-2)
	return memorize[n]
}

func calculateFibonacci(n int) int {
	memorize := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		memorize[i] = -1
	}

	return calculateFibonacciRecurseive(memorize, n)

}

func main() {
	fmt.Println(calculateFibonacci(5))
	fmt.Println(calculateFibonacci(6))
	fmt.Println(calculateFibonacci(7))

}
