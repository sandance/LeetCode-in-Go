package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, word := range strs {
		s := strings.Split(word, "")
		sort.Strings(s)
		key := strings.Join(s, "")
		anagramMap[key] = append(anagramMap[key], word)
	}

	result := [][]string{}
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func main() {
	words := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	fmt.Println(groupAnagrams(words))
}
