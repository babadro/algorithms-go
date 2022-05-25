package _018_4sum

import "sort"

// tptl. passed. medium
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			arrays := searchPairs(nums, i, j, target)

			res = append(res, arrays...)
		}

	}

	return res
}

func searchPairs(nums []int, i, j, target int) [][]int {
	var res [][]int
	l, r := j+1, len(nums)-1
	for l < r {
		sum := nums[i] + nums[j] + nums[l] + nums[r]
		if sum == target {
			res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})

			l++
			r--
			for ; l < r && nums[l] == nums[l-1]; l++ {
			}

			for ; l < r && nums[r] == nums[r+1]; r-- {
			}

		} else if sum < target {
			l++
		} else {
			r--
		}
	}

	return res
}
