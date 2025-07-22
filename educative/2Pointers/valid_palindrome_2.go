package main

import (
	"fmt"
)

/*
Write a function that takes a string as input and checks whether it can be a valid palindrome by removing at most one character from it.


Algorithm:
1.
*/

func validPalindrome(s string) bool {

	isPalindrome := func(str string, left, right int) bool {
		for left < right {
			if str[left] != str[right] {
				return false
			}
			left++
			right--
		}
		return true
	}

	left, right := 0, len(s)-1
	for left <= right {
		if s[left] != s[right] {
			return isPalindrome(s, left, right-1) || isPalindrome(s, left+1, right)
		}
		left++
		right--
	}

	return true
}

func main() {
	inputs := map[string]bool{
		"hello world": false,
		"ABCEBA":      true,
		"DEEAD":       true,
		"abca":        true,
		"racecar":     true,
		"abcdefdba":   false,
	}

	for input, expected := range inputs {
		actual := validPalindrome(input)
		status := "❌ FAIL"
		if actual == expected {
			status = "✅ PASS"
		}
		fmt.Printf("Input: %q | Expected: %v | Got: %v | %s\n", input, expected, actual, status)
	}
}
