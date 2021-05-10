package _1856_maximum_subarray_min_product

// passed. Medium (hard for me). Easy to understand, but very slow. tptl
// https://leetcode.com/problems/maximum-subarray-min-product/discuss/1199635/Java-Simple-and-Easy-understood
func maxSumMinProduct(nums []int) int {
	res := 0

	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			temp := sum(nums, i)

			if t2 := temp * nums[i]; t2 > res {
				res = t2
			}
		}
	}

	return res % 1_000_000_007
}

func sum(nums []int, cur int) int {
	res, val := nums[cur], nums[cur]

	for i := cur - 1; i >= 0 && nums[i] >= val; i-- {
		res += nums[i]
	}

	for i := cur + 1; i < len(nums) && nums[i] >= val; i++ {
		res += nums[i]
	}

	return res
}
