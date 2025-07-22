package main

import "container/heap"

func reorganizeString(str string) string {
	charFreq := make(map[rune]int)
	for _, char := range []rune(str) {
		charFreq[char] += 1
	}

	maxHeap := newMaxHeap()

	for char, freq := range charFreq {
		heap.Push(maxHeap, Set{
			char:  char,
			count: freq,
		})

	}

	prevChar := ""
	prevFreq := 0
	// Replace this placeholder return statement with your code
	return ""
}
