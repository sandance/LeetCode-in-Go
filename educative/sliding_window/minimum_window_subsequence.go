package main

import (
	"fmt"
)

func minWindow(str1 string, str2 string) string {

	charFreq := make(map[byte]int)
	for i := 0; i < len(str2); i++ {
		charFreq[str2[i]]++
	}

	minLen := len(str1) + 1
	windowStart := 0
	substrStart := 0
	matched := 0

	//fmt.Println(charFreq)

	for windowEnd := 0; windowEnd < len(str1); windowEnd++ {
		rightChar := str1[windowEnd]

		if _, exists := charFreq[rightChar]; exists {
			charFreq[rightChar]--
			if charFreq[rightChar] >= 0 {
				matched++
			}
		}

		for matched == len(str2) {
			if currentLen := windowEnd - windowStart + 1; currentLen < minLen {
				minLen = currentLen
				substrStart = windowStart
			}

			leftChar := str1[windowStart]
			windowStart++

			if _, exists := charFreq[leftChar]; exists {
				if charFreq[leftChar] == 0 {
					matched--
				}
				charFreq[leftChar]++
			}
		}
	}

	if minLen > len(str1) {
		return ""
	}
	// Replace this placeholder return statement with your code

	return str1[substrStart : substrStart+minLen]
}

func main() {
	fmt.Println(minWindow("abcdebdde", "bde"))
	//fmt.Println(minWindow("fgrqsqsnodwmxzkzxwqegkndaa", "kzed"))
	//fmt.Println(minWindow("michmznaitnjdnjkdsnmichmznait", "michmznait"))
	//fmt.Println(minWindow("afgegrwgwga", "aa"))
	//fmt.Println(minWindow("abababa", "ba"))
	//fmt.Println(minWindow("zxcvnhss", "css"))
}
