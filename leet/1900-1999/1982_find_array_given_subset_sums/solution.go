package _982_find_array_given_subset_sums

import "sort"

// java version in comments: https://leetcode.com/problems/find-array-given-subset-sums/discuss/1419206/Python3-O(NlogN)-solution-beats-100-of-submissions-in-runtime
func recoverArray(n int, sums []int) []int {
	sumList := make([]int, len(sums))
	copy(sumList, sums)
	sort.Slice(sumList, func(i, j int) bool {
		return sumList[i] > sumList[j]
	})

	res := make([]int, n)
	i := 0
	for len(sumList) > 2 {
		var array1, array2 []int
		num := sumList[0] - sumList[1]
		counts := getCountMap(sumList)
		for _, elm := range sumList {
			if _, ok := counts[elm]; ok {
				array2 = append(array2, elm)
				array1 = append(array1, elm-num)
				remove(counts, elm)
				remove(counts, elm-num)
			}
		}

		index := indexOf(array2, num)
		if index != -1 && array1[index] == 0 {
			res[i] = num
			i++
			sumList = array1
			continue
		}

		res[i] = -num
		i++
		sumList = array2
	}

	if sumList[0] == 0 {
		res[i] = sumList[1]
	} else {
		res[i] = sumList[0]
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
