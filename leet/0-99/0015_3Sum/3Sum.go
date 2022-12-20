package _015_3Sum

import "sort"

// tptl
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		findTriplet(nums[i], nums[i+1:], &res)
	}

	return res
}

func findTriplet(num int, nums []int, triplets *[][]int) {
	left, right := 0, len(nums)-1
	for left < right {
		if left > 0 && left < right && nums[left] == nums[left-1] {
			left++
			continue
		}

		if right < len(nums)-1 && right > left && nums[right] == nums[right+1] {
			right--
			continue
		}

		sum := nums[left] + nums[right] + num

		if sum > 0 {
			right--
		} else if sum < 0 {
			left++
		} else {
			*triplets = append(*triplets, []int{num, nums[left], nums[right]})

			left++
			right--
		}
	}
}
