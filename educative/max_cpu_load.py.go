package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Job struct {
	start    int
	end      int
	cpu_load int
}

// minheap for jobs based on end time
type JobsHeap []*Job

func (h JobsHeap) Len() int           { return len(h) }
func (h JobsHeap) Less(i, j int) bool { return h[i].start < h[j].start }
func (h JobsHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *JobsHeap) Push(x interface{}) {
	*h = append(*h, x.(*Job))
}

func (h *JobsHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func findMaxCPULoad(jobs []*Job) int {
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].start < jobs[j].start
	})

	minHeap := &JobsHeap{}
	heap.Init(minHeap)
	currentCPULoad := 0
	maxCPULoad := 0
	for _, job := range jobs {
		if minHeap.Len() > 0 && job.start >= (*minHeap)[0].end {
			currentCPULoad -= (*minHeap)[0].cpu_load
			heap.Pop(minHeap)
		}

		heap.Push(minHeap, job)
		currentCPULoad += job.cpu_load

		maxCPULoad = max(maxCPULoad, currentCPULoad)
	}
	return maxCPULoad
}

func main() {
	fmt.Println("Maximum CPU load at any time:", findMaxCPULoad([]*Job{
		{1, 4, 3}, {2, 5, 4}, {7, 9, 6},
	}))
	fmt.Println("Maximum CPU load at any time:", findMaxCPULoad([]*Job{
		{6, 7, 10}, {2, 4, 11}, {8, 12, 15},
	}))
	fmt.Println("Maximum CPU load at any time:", findMaxCPULoad([]*Job{
		{1, 4, 2}, {2, 4, 1}, {3, 6, 5},
	}))
}
