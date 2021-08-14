package _001_two_sum

import "sort"

// passed. mine. verbose and slow.
func twoSum(nums []int, target int) []int {
	idxes := make([]int, len(nums))
	for i := range idxes {
		idxes[i] = i
	}

	sort.Slice(idxes, func(i, j int) bool {
		return nums[idxes[i]] < nums[idxes[j]]
	})

	for _, idx1 := range idxes {
		num2 := target - nums[idx1]
		left, right := 0, len(idxes)-1
		var mid int
		for left <= right {
			mid = (left + right) / 2
			if nums[idxes[mid]] < num2 {
				left = mid + 1
			} else if nums[idxes[mid]] > num2 {
				right = mid - 1
			} else {
				if idx1 == idxes[mid] {
					if mid+1 < len(idxes) && nums[idxes[mid+1]] == nums[idxes[mid]] {
						return []int{idx1, idxes[mid+1]}
					}

					return []int{idx1, idxes[mid-1]}
				}

				return []int{idx1, idxes[mid]}
			}
		}
	}

	return nil
}
