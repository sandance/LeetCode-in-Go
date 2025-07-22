package main

import (
	"container/heap"
	"fmt"

	"sort"
	"strings"
)

func minimumMachines(tasks [][]int) int {

	// Replace this placeholder return statement with your code
	sort.Slice(tasks, func(a, b int) bool {
		return tasks[a][0] < tasks[b][0]
	})

	machines := &MinHeap{}
	heap.Init(machines)

	minimumMachinesCount := 0

	for _, task := range tasks {
		start, end := task[0], task[1]
		for machines.Len() > 0 && (*machines)[0] <= start {
			heap.Pop(machines)
		}

		heap.Push(machines, end)
		minimumMachinesCount = max(minimumMachinesCount, machines.Len())
	}
	return minimumMachinesCount
}

// Driver code
func main() {
	inputTasksList := [][][]int{
		{{1, 1}, {5, 5}, {8, 8}, {4, 4}, {6, 6}, {10, 10}, {7, 7}},
		{{1, 7}, {1, 7}, {1, 7}, {1, 7}, {1, 7}, {1, 7}},
		{{1, 7}, {8, 13}, {5, 6}, {10, 14}, {6, 7}},
		{{1, 3}, {3, 5}, {5, 9}, {9, 12}, {12, 13}, {13, 16}, {16, 17}},
		{{12, 13}, {13, 15}, {17, 20}, {13, 14}, {19, 21}, {18, 20}},
	}

	for i, tasks := range inputTasksList {
		fmt.Printf("%d.\tTasks: [", i+1)
		for j, task := range tasks {
			fmt.Printf("[%d, %d]", task[0], task[1])
			if j < len(tasks)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")
		fmt.Printf("\tMinimum number of machines: %d\n", minimumMachines(tasks))
		fmt.Println(strings.Repeat("-", 100))
	}
}
