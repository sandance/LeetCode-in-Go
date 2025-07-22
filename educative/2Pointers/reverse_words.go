package main

import (
	"fmt"
	"strings"
)

func replace_sentence(str1 string) string {
	str1 = strings.TrimSpace(str1)
	words := strings.Fields(str1)

	left, right := 0, len(words)-1
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	return strings.Join(words, " ")
}

func main() {
	inputs := []string{
		"Reverse this string",
		" smart",
		" MULTIPLE SPACES",
	}

	for _, input := range inputs {
		fmt.Println("Input: ", input)
		result := replace_sentence(input)
		fmt.Println(result)
	}
}
