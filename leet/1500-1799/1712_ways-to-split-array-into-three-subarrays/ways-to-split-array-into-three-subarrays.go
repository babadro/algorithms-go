package _1712_ways_to_split_array_into_three_subarrays

import (
	"fmt"
	"sort"
)

// todo 1
// https://leetcode.com/problems/ways-to-split-array-into-three-subarrays/discuss/1005689/Java-O(N)-sliding-window
// https://leetcode.com/problems/ways-to-split-array-into-three-subarrays/discuss/999257/C%2B%2BJavaPython-O(n)-with-picture
func waysToSplit2(nums []int) int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}

	if nums[n-1] == 0 {
		k := n - 2
		return (1 + k) * k / 2
	}

	counter := 0
	for i := 0; i < n-2; i++ {

		tmp := nums[i+1 : n-1]
		k := len(tmp)
		start := sort.Search(k, func(j int) bool {
			return tmp[j]-nums[i] >= nums[i]
		})
		start += i + 1

		tmp2 := nums[start+1 : n]
		finish := sort.Search(len(tmp2), func(j int) bool {
			return nums[j]-nums[i] > nums[n-1]-nums[j]
		})
		if finish >= len(tmp2) {
			continue
		}

		count := finish - start
		if count <= 0 {
			continue
		}

		counter += count
	}

	return counter % 1_000_000_007
}

func waysToSplit(nums []int) int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}

	counter := 0
	for i := 0; i < n-2; i++ {
		start, finish := -1, n-1
		for j := i + 1; j < n-1; j++ {
			if start == -1 {
				if nums[j]-nums[i] >= nums[i] {
					start = j
				}
			} else if nums[j]-nums[i] > nums[n-1]-nums[j] {
				finish = j
				break
			}
		}

		if start == -1 || nums[n-1]-nums[finish-1] < nums[finish-1]-nums[i] {
			continue
		}

		counter += finish - start

		if counter == 227 {

			fmt.Printf("i=%d, start=%d, finish=%d\n", i, start, finish)
		}
	}

	return counter % 1_000_000_007
}

//if nums[i] >= nums[n-1]-nums[i] {
//	break
//}
//tmp := nums[i+1:n-2]
//nTmp := len(tmp)
//
//start := sort.Search(nTmp, func(j int) bool {
//	return tmp[j]-nums[i] >= nums[i]
//})
//
//finish := sort.Search(nTmp, func(j int) bool {
//	return tmp[j]-nums[i] > nums[n-1]-tmp[j]
//})
