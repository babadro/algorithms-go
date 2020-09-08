package _152_maximum_product_subarray

// TODO 1
// caterpillar?
func maxProduct(nums []int) int {
	i, j, maxProduct, currProduct := 0, 1, nums[0], nums[0]
	for i < len(nums) || j < len(nums) {
		for j < len(nums) && nums[j] != 0 {
			if currProduct *= nums[j]; currProduct > maxProduct {
				0
			}
		}
	}

	return maxProduct
}
