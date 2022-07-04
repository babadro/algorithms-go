// TODO 2 look for java solution in https://leetcode.com/problems/maximum-product-subarray/discuss/183483/JavaC%2B%2BPython-it-can-be-more-simple
package _152_maximum_product_subarray

// tptl passed. best solution. not mine
func maxProduct(nums []int) int {
	res, curMin, curMax := nums[0], nums[0], nums[0]
	for _, v := range nums[1:] {
		if v < 0 {
			curMin, curMax = curMax, curMin
		}

		curMin = min(curMin*v, v)
		curMax = max(curMax*v, v)

		res = max(res, curMax)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// caterpillar 100% 77%
func maxProduct2(nums []int) int {
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
