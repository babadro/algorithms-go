package _7_employee_free_time

import "container/heap"

type interval struct {
	empIDx, intvIDx int
}

func findEmployeeFreeTime(schedule [][][]int) [][]int {
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