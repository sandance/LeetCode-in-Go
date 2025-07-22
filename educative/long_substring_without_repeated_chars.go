package main

import (
	"fmt"
	"github.com/golang-collections/collections/set"
)

func findMaxRepeatedString(str1 string) int {
	charSet := set.Set{}
	window_start := 0
	max_length := 0

	for window_end, rightChar := range str1 {
		for charSet.Has(rightChar) {
			leftChar := str1[window_start]
			charSet.Remove(leftChar)
			window_start += 1
		}

		charSet.Insert(rightChar)
		max_length = max(max_length, window_end-window_start+1)

	}
	return max_length
}

func main() {
	fmt.Println(findMaxRepeatedString("aabccbb"))
}
