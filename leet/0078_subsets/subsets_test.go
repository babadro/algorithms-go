package _078_subsets

import "testing"

func TestSubsets(t *testing.T) {
	nums := make([]int, 30)
	for i := 0; i < 30; i++ {
		nums[i] = i
	}
	res := subsets(nums)
	t.Log(res)
	t.Log(len(res))
}
