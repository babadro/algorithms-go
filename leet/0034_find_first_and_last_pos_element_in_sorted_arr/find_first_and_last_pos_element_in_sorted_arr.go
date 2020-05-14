package _034_find_first_and_last_pos_element_in_sorted_arr

import "sort"

// TODO 2 implement better solution. Implement custom binary search maybe?
// Runtime: 8 ms, faster than 79.40% Memory Usage: 4.1 MB, less than 50.00%
// Еще проблема может быть в том, что дается очень длинный кусок повторяющихся значений - и простой
// перебор по ним дает плохой результат - нужно бинарным поиском найти и конец сегмента
// Хотя все равно заменил на поиск и второго конца - результат тот же 8 мс и 4.1 мб
func searchRange(nums []int, target int) []int {
	l := len(nums)
	start := sort.Search(l, func(i int) bool {
		return nums[i] >= target
	})
	if start >= l || nums[start] != target {
		return []int{-1, -1}
	}
	nums = nums[start:]
	end := sort.Search(len(nums), func(i int) bool {
		return nums[i] != target
	})
	end += start - 1
	return []int{start, end}
}
