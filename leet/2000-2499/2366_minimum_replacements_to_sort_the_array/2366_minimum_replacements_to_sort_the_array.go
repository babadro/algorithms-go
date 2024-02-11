package _2366_minimum_replacements_to_sort_the_array

func minimumReplacement(nums []int) int64 {
	var res, cur, prev int

	if len(nums) > 0 {
		prev = nums[len(nums)-1]
	}

	for i := len(nums) - 2; i >= 0; {
		cur = nums[i]

		for cur > prev {
			res++

			remainder := cur / 2

			newPrev := cur/2 + remainder

			cur /= 2
		}

		prev = cur
		i--
	}

	return int64(res)
}
