package _334_increasing_triplet_subsequence

// Doesn't work
func increasingTripletFail(nums []int) bool {
	i := 0

Loop:
	for i < len(nums) {
		for idx := i + 1; idx < i+3; idx++ {
			if idx == len(nums) {
				return false
			}
			if nums[idx] <= nums[idx-1] {
				i = idx
				continue Loop
			}
		}

		return true
	}

	return false
}
