package _1712_ways_to_split_array_into_three_subarrays

import (
	"fmt"
	"sort"
)

// todo 1
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

		start := sort.Search(n-3-i, func(j int) bool {
			return nums[j+i+1]-nums[i] >= nums[i]
		})
		start += i + 1

		finish := sort.Search(n-start, func(j int) bool {
			return nums[j+start]-nums[i] > nums[n-1]-nums[j]
		})
		finish += start

		count := finish - start
		if count <= 0 || finish >= n || start <= i {
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
