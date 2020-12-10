package _228_summary_ranges

import (
	"fmt"
	"strconv"
)

// passed. best solution. tptl
func summaryRanges2(nums []int) (res []string) {
	for i, j := 0, 0; j < len(nums); j++ {
		if j < len(nums)-1 && nums[j+1]-nums[j] == 1 {
			continue
		}

		if nums[i] == nums[j] {
			res = append(res, strconv.Itoa(nums[i]))
		} else {
			res = append(res, fmt.Sprintf("%s->%s", strconv.Itoa(nums[i]), strconv.Itoa(nums[j])))
		}

		i = j + 1
	}

	return res
}

// passed
func summaryRanges(nums []int) (res []string) {
	if len(nums) == 0 {
		return res
	}

	start, finish := nums[0], nums[0]
	for i := 1; i <= len(nums); i++ {
		if i == len(nums) || nums[i]-nums[i-1] > 1 {
			if start == finish {
				res = append(res, strconv.Itoa(start))
			} else {
				res = append(res, fmt.Sprintf("%s->%s", strconv.Itoa(start), strconv.Itoa(finish)))
			}

			if i < len(nums) {
				start, finish = nums[i], nums[i]
			}
		} else {
			finish = nums[i]
		}
	}

	return res
}
