package _347_top_k_frequent_elements

import "sort"

// bruteforce 93% 5%
func topKFrequent(nums []int, k int) []int {
	set := make(map[int]int, len(nums))

	for _, v := range nums {
		set[v]++
	}

	arr := make([][2]int, 0, len(set))

	for value, frequency := range set {
		arr = append(arr, [2]int{value, frequency})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})

	res := make([]int, k)

	for i := 0; i < k; i++ {
		res[i] = arr[i][0]
	}

	return res
}

// TODO 2 short solution (BucketSort)
// https://leetcode.com/problems/top-k-frequent-elements/discuss/371250/Go-MinHeap-and-Bucket-sort-solution
