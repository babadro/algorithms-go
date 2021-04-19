package _347_top_k_frequent_elements

import "math/rand"

// Approach 2 - QuickSelect https://leetcode.com/problems/top-k-frequent-elements/solution/
// TODO 2 doesn't work fix func and tests.
func topKFrequent3(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	n := len(count)
	unique := make([]int, 0, n)
	for key := range count {
		unique = append(unique, key)
	}

	quickSelect3(0, n-1, n-k, count, unique)

	return unique //unique[n-k:n]

}

func partition3(left, right, pivotIdx int, count map[int]int, unique []int) int {
	pivotFrequency := count[pivotIdx]
	unique[pivotIdx], unique[right] = unique[right], unique[pivotIdx]

	storeIdx := left

	for i := left; i <= right; i++ {
		if count[unique[i]] < pivotFrequency {
			unique[storeIdx], unique[i] = unique[i], unique[storeIdx]
			storeIdx++
		}
	}

	unique[storeIdx], unique[right] = unique[right], unique[storeIdx]

	return storeIdx
}

func quickSelect3(left, right, kSmallest int, count map[int]int, unique []int) {
	if left == right {
		return
	}

	pivotIdx := rand.Intn(right - left)

	pivotIdx = partition3(left, right, pivotIdx, count, unique)

	if kSmallest == pivotIdx {
		return
	} else if kSmallest < pivotIdx {
		quickSelect3(left, pivotIdx-1, kSmallest, count, unique)
	} else {
		quickSelect3(pivotIdx+1, right, kSmallest, count, unique)
	}
}

// Approach 2 - QuickSelect
// Doesn't work: fatal error: out of memory allocating heap arena metadata
// https://leetcode.com/problems/top-k-frequent-elements/discuss/270988/Javascript-96-Quick-Select-Easy-implementation-to-remember
func topKFrequent2(nums []int, k int) []int {
	if len(nums) == k {
		return nums
	}

	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	uniq := make([]int, 0, len(count))
	for key := range count {
		uniq = append(uniq, key)
	}

	quickSelect(uniq, count, 0, len(uniq)-1, k)

	return uniq[:k]
}

func quickSelect(nums []int, m map[int]int, start, end, k int) {
	pivot := nums[(start+end)/2]
	low, high := start, end

	for low <= high {
		for m[nums[low]] > m[pivot] && low <= high {
			low++
		}
		for m[nums[high]] < m[pivot] && low <= high {
			high--
		}

		if low <= high {
			nums[low], nums[high] = nums[high], nums[low]
			low++
			high--
		}
	}

	if low <= k {
		quickSelect(nums, m, low, end, k)
	}
	if high >= k {
		quickSelect(nums, m, start, high, k)
	}
}
