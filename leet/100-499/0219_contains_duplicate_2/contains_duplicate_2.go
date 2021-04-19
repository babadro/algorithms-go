package _219_contains_duplicate_2

import "sort"

// passed. fast but not best for memory using. good choice because short.
func containsNearbyDuplicate2(nums []int, k int) bool {
	m := make(map[int]int)
	for i, num := range nums {
		if idx, ok := m[num]; ok && i-idx <= k {
			return true
		}

		m[num] = i
	}

	return false
}

// passed. tptl. best solution
func containsNearbyDuplicate(nums []int, k int) bool {
	arr := make([][2]int, len(nums))

	for i, num := range nums {
		arr[i][0], arr[i][1] = num, i
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			return arr[i][1] < arr[j][1]
		}

		return arr[i][0] < arr[j][0]
	})

	for i := 0; i < len(arr)-1; i++ {
		if arr[i][0] != arr[i+1][0] {
			continue
		}

		idx1, idx2 := arr[i][1], arr[i+1][1]
		if idx1 < idx2 {
			idx1, idx2 = idx2, idx1
		}

		if idx1-idx2 <= k {
			return true
		}
	}

	return false
}
