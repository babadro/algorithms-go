package _0759_employee_free_time

import (
	"sort"
)

type Interval struct {
	Start int
	End   int
}

// bnsrg passed hard
func employeeFreeTime(schedule [][]*Interval) []*Interval {
	var intervals []*Interval

	for _, employee := range schedule {
		intervals = append(intervals, employee...)
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	end := intervals[0].End

	var res []*Interval

	for _, interval := range intervals[1:] {
		if interval.Start > end {
			res = append(res, &Interval{
				Start: end,
				End:   interval.Start,
			})

			end = interval.End

			continue
		}

		end = max(end, interval.End)
	}

	return res
}
