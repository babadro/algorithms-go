package _152_maximum_product_subarray

// caterpillar 100% 77%
func maxProduct(nums []int) int {
	front, maxProduct := 0, nums[0]
	for front < len(nums) {
		if nums[front] == 0 {
			if maxProduct < 0 {
				maxProduct = 0
			}
			front++
			continue
		}

		end := front
		currProduct := 1
		for ; front < len(nums) && nums[front] != 0; front++ {
			if currProduct *= nums[front]; currProduct > maxProduct {
				maxProduct = currProduct
			}
		}

		for ; currProduct < 0 && end < front-1; end++ {
			if currProduct /= nums[end]; currProduct > maxProduct {
				maxProduct = currProduct
			}
		}
	}

	return maxProduct
}

// TODO 2 look for java solution in https://leetcode.com/problems/maximum-product-subarray/discuss/183483/JavaC%2B%2BPython-it-can-be-more-simple

// TODO 2 understand
// https://leetcode.com/problems/maximum-product-subarray/discuss/302805/Simple-code-in-Go-0ms
func maxProduct2(nums []int) int {
	tmp := nums[0]
	max := nums[0]
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			max, min = min, max
		}
		max *= nums[i]
		if max < nums[i] {
			max = nums[i]
		}
		min *= nums[i]
		if min > nums[i] {
			min = nums[i]
		}
		if max > tmp {
			tmp = max
		}
	}
	return tmp
}
