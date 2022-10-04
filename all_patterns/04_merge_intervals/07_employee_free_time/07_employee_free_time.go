package _7_employee_free_time

import (
	"math"
	"sort"

	"github.com/babadro/algorithms-go/utils"
)

// leetcode premium https://leetcode.com/problems/employee-free-time/
// todo  doesn't work
func findEmployeeFreeTime(schedule [][]int) [][]int {
	minTime, maxTime := math.MaxInt64, math.MinInt64
	for _, emp := range schedule {
		minTime = utils.Min(minTime, emp[0])
		maxTime = utils.Max(maxTime, emp[len(emp)-1])
	}

	var freeIntervals [][]int
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

	sort.Slice(freeIntervals, func(i, j int) bool {
		return freeIntervals[i][0] < freeIntervals[j][0]
	})

	counter := 0
	empCount := len(schedule)
	var res [][]int
	var cur []int
	for _, intv := range freeIntervals {
		if counter == 0 || cur[1] <= intv[0] {
			cur = intv
			counter = 1
		} else {
			cur[0] = intv[0]
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
