package _7_employee_free_time

import (
	"container/heap"
	"math"

	"github.com/babadro/algorithms-go/utils"
)

// leetcode premium https://leetcode.com/problems/employee-free-time/
// todo  check solution with more test cases
func findEmployeeFreeTime(schedule [][]int) [][]int {
	minTime, maxTime := math.MaxInt64, math.MinInt64
	for _, emp := range schedule {
		minTime = utils.Min(minTime, emp[0])
		maxTime = utils.Max(maxTime, emp[len(emp)-1])
	}

	var freeIntervals minHeap
	for _, emp := range schedule {
		if minTime < emp[0] {
			freeIntervals = append(freeIntervals, []int{minTime, emp[0]})
		}

		if maxTime > emp[len(emp)-1] {
			freeIntervals = append(freeIntervals, []int{emp[len(emp)-1], maxTime})
		}

		for i := 1; i < len(emp)-1; i += 2 {
			freeIntervals = append(freeIntervals, []int{emp[i], emp[i+1]})
		}
	}

	heap.Init(&freeIntervals)

	counter := 0
	empCount := len(schedule)
	var res [][]int
	var cur []int
	for len(freeIntervals) > 0 {
		intv := freeIntervals[0]
		_ = heap.Pop(&freeIntervals)

		if counter == 0 || cur[1] <= intv[0] {
			cur = intv
			counter = 1
		} else {
			cur[0] = intv[0]
			if cur[1] < intv[1] {
				heap.Push(&freeIntervals, []int{cur[1], intv[1]})
			} else if cur[1] > intv[1] {
				heap.Push(&freeIntervals, []int{intv[1], cur[1]})
			}

			cur[1] = utils.Min(cur[1], intv[1])
			counter++
		}

		if counter == empCount {
			res = append(res, cur)
			counter = 0
		}
	}

	return res
}

type minHeap [][]int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(v interface{}) { *h = append(*h, v.([]int)) }
func (h *minHeap) Pop() interface{} {
	last := len(*h) - 1
	res := (*h)[last]
	*h = (*h)[:last]

	return res
}
