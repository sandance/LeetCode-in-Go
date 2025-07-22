package main

import (
	"fmt"
	"strconv"
)

func add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	//if b == 0 {
	//	return 0, errors.New("divide by zero")
	//}
	return a / b
}

var opMap = map[string]func(int, int) int{
	"+": add,
	"-": Subtract,
	"*": Multiply,
	"/": Divide,
}

func main() {
	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"4", "*", "5"},
		[]string{"6", "/", "7"},
		[]string{"8", "-", "9"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("Wrong number of arguments passed to function", expression)
			continue
		}

		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("Unsupported operation", op)
			continue
		}

		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
		}

		result := opFunc((p1), (p2))
		fmt.Println(result)
	}
}
