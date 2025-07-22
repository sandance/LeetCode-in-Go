package main

import (
	"fmt"
)

func longestSubstringWithKDistinct(str string, k int) int {
	charFreq := make(map[byte]int)
	windowStart := 0
	maxLength := 0

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]
		charFreq[rightChar] += 1

		for len(charFreq) > k {
			leftChar := str[windowStart]
			charFreq[leftChar] -= 1
			if charFreq[leftChar] == 0 {
				delete(charFreq, leftChar)
			}
			windowStart += 1
		}
		maxLength = max(maxLength, windowEnd-windowStart+1)

	}
	return maxLength
}

func main() {
	fmt.Println(longestSubstringWithKDistinct("abaaci", 2))
	fmt.Println(longestSubstringWithKDistinct("araacc", 2))
	fmt.Println(longestSubstringWithKDistinct("cbbebi", 3))

	fmt.Println(longestSubstringWithKDistinct("abaaci", 2))
}
