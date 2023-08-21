package _0759_employee_free_time

import "math"

type Interval struct {
	Start int
	End   int
}

type iterator struct {
	arrIDx int
	next   int
}

const inf = math.MaxInt64

// bnsrg todo 1
func employeeFreeTime(schedule [][]*Interval) []*Interval {
	min, idx := math.MaxInt64, -1

	for i, in := range schedule {
		if end := in[0].End; end < min {
			min, idx = end, i
		}
	}

	schedule[idx], schedule[0] = schedule[0], schedule[idx]

	iterators := make([]iterator, len(schedule))
	for i := range iterators {
		iterators[i].arrIDx = i
	}

	left, right := schedule[0][0].End, inf
	if len(schedule) > 1 {
		right = schedule[0][1].Start
	}

	for i := 1; ; i = (i + 1) % len(iterators) {
		it := &iterators[i]
		interval := schedule[it.arrIDx][it.next]

		for ; interval.End <= left; it.next++ {
			if it.next == len(schedule[it.arrIDx]) {
				break
			}

		}

	}

}
