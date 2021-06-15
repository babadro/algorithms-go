package _033_search_in_rotated_sorted_array

// задумка была сделать такую модификацию алгоритма: элементы в массиве могут быть не уникальными.
// Вернуть самый левый индекс искомого числа, если оно есть в массиве, или -1, если нет.
// passed. tptl (?)
func searchNotUnique(nums []int, target int) int {
	n := len(nums)
	to := n
	for i := n - 1; i > 0 && nums[i] == nums[0]; i-- {
		to--
	}

	startRight := findStartRight(nums, 0, to)

	if target < nums[0] {
		return binarySearchTarget(nums, startRight, n, target)
	}

	return binarySearchTarget(nums, 0, startRight, target)
}

// find the start of the right part of array
func findStartRight(nums []int, from, to int) int {
	var res int
	first, last := nums[0], nums[to-1]

	for from < to {
		mid := (from + to) / 2
		curr := nums[mid]
		if curr > last {
			from = mid + 1
		} else if curr < first {
			res, to = mid, mid
		} else {
			return len(nums)
		}
	}

	return res
}

// find leftmost index where nums[index] == target. If not exists return -1
func binarySearchTarget(nums []int, left, right, target int) int {
	var res int
	for left < right {
		mid := (left + right) / 2
		curr := nums[mid]
		if target > curr {
			left = mid + 1
		} else if target <= curr {
			res, right = mid, mid
		}
	}

	if nums[res] == target {
		return res
	}

	return -1
}
