package _6_maximum_cpu_load

import (
	"container/heap"
	"sort"

	"github.com/babadro/algorithms-go/utils"
)

// no leetcode example
// Given an array of jobs with different time requirements,
// where each job consists of start time, end time and CPU load.
// The task is to find the maximum CPU load at any time if all jobs
// are running on the same machine.
func findMaxCPULoad(jobs [][3]int) int {
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})

	var h minHeap
	var maxLoad, curLoad int
	for _, job := range jobs {
		now := job[0]
		for h.Len() > 0 && h[0][1] <= now {
			curLoad -= h[0][2]
			heap.Pop(&h)
		}

		heap.Push(&h, job)
		curLoad += job[2]

		maxLoad = utils.Max(maxLoad, curLoad)
	}

	return maxLoad
}

type minHeap [][3]int

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *minHeap) Push(v interface{}) {
	*h = append(*h, v.([3]int))
}
func (h *minHeap) Pop() interface{} {
	last := len(*h) - 1
	res := (*h)[last]
	*h = (*h)[:last]

	return res
}
