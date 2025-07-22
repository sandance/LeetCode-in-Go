package main

func longestRepeatingCharacterReplacement(s string, k int) int {
	// Replace this placeholder return statement with your code
	charFreq := make(map[byte]int, 0)
	maxRepeatedCharCount := 0
	windowStart := 0
	maxLen := 0

	for windowEnd := 0; windowEnd < len(s); windowEnd++ {
		rightChar := s[windowEnd]
		charFreq[rightChar]++

		maxRepeatedCharCount = max(maxRepeatedCharCount, charFreq[rightChar])

		for windowEnd-windowStart+1-maxRepeatedCharCount > k {
			leftChar := s[windowStart]
			charFreq[leftChar]--
			windowStart++
		}
		maxLen = max(maxLen, windowEnd-windowStart+1)
	}
	return maxLen
}
