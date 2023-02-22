package _7_employee_free_time

import (
	"container/heap"
	"sort"
)

type interval struct {
	empIDx, intvIDx int
}

// leetcode premium: https://leetcode.com/problems/employee-free-time/
// We are given a list scheduleof employees, which represents the working time for each employee.
// Each employee has a list of non-overlapping Intervals, and these intervals are in sorted order.
// Return the list of finite intervals representing common, positive-length free time for all employees,
// also in sorted order.

// tptl best solution
// https://aaronice.gitbook.io/lintcode/sweep-line/employee-free-time
func findEmployeeFreeTime(schedule [][][]int) [][]int {
	var result, timeLine [][]int

	for _, emp := range schedule {
		timeLine = append(timeLine, emp...)
	}

	sort.Slice(timeLine, func(i, j int) bool {
		return timeLine[i][0] < timeLine[j][0]
	})

	temp := timeLine[0]
	for _, tInterval := range timeLine {
		if temp[1] < tInterval[0] {
			result = append(result, []int{temp[1], tInterval[0]})
			temp = tInterval
		} else if temp[1] < tInterval[1] {
			temp = tInterval
		}
	}

	return result
}

// overcomplicated
func findEmployeeFreeTime2(schedule [][][]int) [][]int {
	var result [][]int

	// PriorityQueue to store one interval from each employee
	h := minHeap{schedule: schedule}

	// insert the first interval of each employee to the queue
	for i := range schedule {
		h.intervals = append(h.intervals, interval{i, 0})
	}

	p := h.intervals[0]
	prev := schedule[p.empIDx][p.intvIDx]
	for h.Len() > 0 {
		t := heap.Pop(&h).(interval)
		top := schedule[t.empIDx][t.intvIDx]

		// if previousInterval is not overlapping with the next interval, insert a free interval
		if prev[1] < top[0] {
			result = append(result, []int{prev[1], top[0]})
			prev = top
		} else { // overlapping intervals, update the previousInterval if needed
			if prev[1] < top[1] {
				prev = top
			}
		}

		empSchedule := schedule[t.empIDx]
		if len(empSchedule) > t.intvIDx+1 {
			heap.Push(&h, interval{t.empIDx, t.intvIDx + 1})
		}
	}

	return result
}

type minHeap struct {
	intervals []interval
	schedule  [][][]int
}

func (h minHeap) Len() int { return len(h.intervals) }
func (h minHeap) Less(i, j int) bool {
	intI, intJ := h.intervals[i], h.intervals[j]
	return h.schedule[intI.empIDx][intI.intvIDx][0] < h.schedule[intJ.empIDx][intJ.intvIDx][0]
}
func (h minHeap) Swap(i, j int)       { h.intervals[i], h.intervals[j] = h.intervals[j], h.intervals[i] }
func (h *minHeap) Push(v interface{}) { h.intervals = append(h.intervals, v.(interval)) }
func (h *minHeap) Pop() interface{} {
	last := len(h.intervals) - 1
	res := h.intervals[last]
	h.intervals = h.intervals[:last]

	return res
}
