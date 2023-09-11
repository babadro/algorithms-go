package _0759_employee_free_time

import "math"

type Interval struct {
	Start int
	End   int
}

const inf = math.MaxInt64

// bnsrg todo 1
func employeeFreeTime(schedule [][]*Interval) []*Interval {
	minStart, idx := math.MaxInt64, -1

	for i, in := range schedule {
		if end := in[0].End; end < minStart {
			minStart, idx = end, i
		}
	}

	schedule[idx], schedule[0] = schedule[0], schedule[idx]

	iterators := make([]int, len(schedule))

	left, right := schedule[0][0].End, inf
	if len(schedule) > 1 {
		right = schedule[0][1].Start
	}

	counter := 0
	for i := 1; ; i = (i + 1) % len(schedule) {
		for ; iterators[i] < len(schedule[i]); iterators[i]++ {
			interval := schedule[i][iterators[i]]

			if interval.End <= left {
				continue
			}

			if interval.Start > left {
				right = min(right, interval.Start)
				break
			}

			left = max(left, interval.End)
		}

	}

}
