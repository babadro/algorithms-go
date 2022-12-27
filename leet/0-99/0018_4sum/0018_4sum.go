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

// slightly alternative approach - just reuse 3sum function
func fourSum2(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		threeSum(nums[i+1:], nums[i], target, &res)
	}

	return res
}

func threeSum(nums []int, num, target int, res *[][]int) {
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		findTriplet(nums[i+1:], num, nums[i], target, res)
	}
}

func findTriplet(nums []int, num1, num2, target int, res *[][]int) {
	left, right := 0, len(nums)-1
	for left < right {
		if left > 0 && nums[left] == nums[left-1] {
			left++
			continue
		}

		if right < len(nums)-1 && nums[right] == nums[right+1] {
			right--
			continue
		}

		if sum := num1 + num2 + nums[left] + nums[right]; sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			*res = append(*res, []int{num1, num2, nums[left], nums[right]})

			left++
			right--
		}
	}
}
