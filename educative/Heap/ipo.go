package main

import (
	"container/heap"
	"fmt"
)

func maxCapital(c int, k int, capitals []int, profits []int) int {
	currentCapital := 2
	profitsMaxHeap := newMaxHeap()

	// Select projects (in the range of current capital) then push them in max heap
	fmt.Printf("\tIdentifying projects we can afford with our current capital of %d:\n", currentCapital)
	for index, value := range profits {
		if len(capitals) != 0 && capitals[index] <= currentCapital {
			heap.Push(profitsMaxHeap, value)
			fmt.Printf("\t\tProfit of project (with capital requirement of %d) pushed on to max heap = %d\n", currentCapital, value)
		}
	}

	return profitsMaxHeap
}
