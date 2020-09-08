package _152_maximum_product_subarray

// TODO 1
// caterpillar?
func maxProduct(nums []int) int {
	i, j, maxProduct, currProduct := 0, 1, nums[0], nums[0]
	for i < len(nums) || j < len(nums) {
		for j < len(nums) && nums[j] != 0 {
			if currProduct == 0 {
				currProduct = 1
			}
			if currProduct *= nums[j]; currProduct > maxProduct {
				maxProduct = currProduct
			}
			j++
		}

		for currProduct < 0 && i < len(nums) && nums[i] != 0 {
			if currProduct /= nums[i]; currProduct > maxProduct {
				maxProduct = currProduct
			}
			i++
		}

		currProduct = 0
		i++
		j++
	}

	return maxProduct
}
