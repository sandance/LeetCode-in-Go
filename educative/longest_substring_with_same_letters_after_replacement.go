package main

import "fmt"

func lengthOflongestSubstringK(str string, k int) int {
	charFreq := make(map[byte]int)
	windowStart := 0
	maximumLength := 0
	maximumRepeatedChar := 0

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]
		charFreq[rightChar] += 1
		maximumRepeatedChar = max(maximumRepeatedChar, charFreq[rightChar])

		for (windowEnd - windowStart + 1 - maximumRepeatedChar) > k {
			leftChar := str[windowStart]
			charFreq[leftChar] -= 1
			windowStart += 1
		}
		maximumLength = max(maximumLength, windowEnd-windowStart+1)
	}
	return maximumLength
}

func main() {
	fmt.Println(lengthOflongestSubstringK("abab", 2))
	fmt.Println(lengthOflongestSubstringK("aabccbb", 2))
	fmt.Println(lengthOflongestSubstringK("abbcb", 1))
}
