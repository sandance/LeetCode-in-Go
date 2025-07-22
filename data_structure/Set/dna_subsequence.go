package main

import (
	"fmt"
	"github.com/golang-collections/collections/set"
)

func findRepeatedDnaSequences(s string) []string {
	result := make([]string, 0)
	seen := set.New()
	k := 10

	for i := 0; i < len(s)-k-1; i++ {
		currStr := s[i : k+i]
		if seen.Has(currStr) {
			result = append(result, currStr)
		}
		seen.Insert(currStr)
	}

	// Replace this placeholder return statement with your code

	return result
}

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Println(findRepeatedDnaSequences(s))
}
