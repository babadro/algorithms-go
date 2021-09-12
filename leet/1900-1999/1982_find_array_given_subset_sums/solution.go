package _982_find_array_given_subset_sums

import "sort"

// java version in comments: https://leetcode.com/problems/find-array-given-subset-sums/discuss/1419206/Python3-O(NlogN)-solution-beats-100-of-submissions-in-runtime
// todo 1 read the explanation
func recoverArray(n int, sums []int) []int {
	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})

	res := make([]int, n)
	i := 0
	for ; len(sums) > 2; i++ {
		var array1, array2 []int
		num := sums[0] - sums[1]
		counts := getCountMap(sums)
		for _, elm := range sums {
			if _, ok := counts[elm]; ok {
				array2 = append(array2, elm)
				array1 = append(array1, elm-num)
				remove(counts, elm)
				remove(counts, elm-num)
			}
		}

		idx := indexOf(array2, num)
		if idx != -1 && array1[idx] == 0 {
			res[i] = num
			sums = array1
		} else {
			res[i] = -num
			sums = array2
		}
	}

	if sums[0] == 0 {
		res[i] = sums[1]
	} else {
		res[i] = sums[0]
	}

	return res
}

func indexOf(nums []int, num int) int {
	for i := range nums {
		if num == nums[i] {
			return i
		}
	}

	return -1
}

func getCountMap(nums []int) map[int]int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	return m
}

func remove(m map[int]int, elm int) {
	if val, ok := m[elm]; ok {
		if val > 1 {
			m[elm]--
		} else {
			delete(m, elm)
		}
	}
}
