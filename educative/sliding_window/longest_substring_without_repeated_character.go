package main

import "fmt"

func findLongestSubstring(str string) int {

	// Replace this placeholder return statement with your code
	charIndex := make(map[byte]int)
	windowStart := 0
	maxWindow := 0

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		char := str[windowEnd]

		if lastIndex, exists := charIndex[char]; exists && lastIndex >= windowStart {
			windowStart = lastIndex + 1
		}
		charIndex[char] = windowEnd
		maxWindow = max(maxWindow, windowEnd-windowStart+1)
	}

	return maxWindow
}

func main() {
	fmt.Println(findLongestSubstring("abcabcbb")) // 3 ("abc")
	fmt.Println(findLongestSubstring("bbbbb"))    // 1 ("b")
	fmt.Println(findLongestSubstring("pwwkew"))   // 3 ("wke")
}
