package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			fmt.Printf("skipping non alphanumeric character %c\n", s[left])
			left++
		}

		for left < right && !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			fmt.Printf("skipping non alphanumeric character %c\n", s[right])
			right--
		}

		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}

func isPalindrome2(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if strings.ToLower(strconv.Itoa(int(s[left]))) != strings.ToLower(strconv.Itoa(int(s[right]))) {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println("Palindrome", isPalindrome2("cbabd"))

	testCases := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		"1A@2!3 23!2@a1",
		"No 'x' in Nixon",
		"12321",
	}

	for _, testCase := range testCases {
		fmt.Printf("\n\tString: %s\n\n", testCase)
		fmt.Println(isPalindrome2(testCase))
		fmt.Println(strings.Repeat("-", 100))
	}
}
