package _015_3Sum

import "sort"

//TODO 2 this is brute force. See hint 3 and if needed solution
func threeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return nil
	}
	sort.Ints(nums)
	resMap := make(map[struct{ a, b, c int }]bool)
	res := make([][]int, 0)
	//length := removeDuplicates(nums)
	//nums = nums[:length]
	//if length < 3 {
	//	return nil
	//}
	tmp := make([]int, 3)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			idx := sort.Search(length, func(k int) bool {
				return nums[k] >= -(nums[i] + nums[j])
			})

			if idx < length && idx != j && idx != i && nums[idx] == -(nums[i]+nums[j]) {
				tmp = tmp[:0]
				tmp = append(tmp, []int{nums[i], nums[j], nums[idx]}...)
				sort.Ints(tmp)
				key := struct{ a, b, c int }{tmp[0], tmp[1], tmp[2]}
				if !resMap[key] {
					res = append(res, []int{nums[i], nums[j], nums[idx]})
					resMap[key] = true
				}
			}
		}
	}

	return res
}

func removeDuplicates(nums []int) int {
	lastPos := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[lastPos] {
			nums[lastPos+1] = nums[i]
			lastPos++
		}
	}
	return lastPos + 1
}
