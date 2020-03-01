package _365_how_many_numbers_are_smaller_than_the_curr_num

import "sort"

type item struct {
	num              int
	idx              int
	countLessNumbers int
}

// naive
func smallerNumbersThanCurrent2(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if nums[i] > nums[j] {
				cnt++
			}
			ans[i] = cnt
		}
	}
	return ans
}

func smallerNumbersThanCurrent(nums []int) []int {
	items := make([]item, len(nums))
	for i, num := range nums {
		items[i].num = num
		items[i].idx = i
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].num < items[j].num
	})

	last := 0
	for i := 0; i < len(items); i++ {
		num := items[i].num
		if num != items[last].num {
			items[i].countLessNumbers = i
			last = i
		} else {
			items[i].countLessNumbers = last
		}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].idx < items[j].idx
	})
	res := make([]int, len(nums))
	for i := range items {
		res[i] = items[i].countLessNumbers
	}
	return res
}
