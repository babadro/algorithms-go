package _665_non_decreasing_array

// tptl. best solution
func checkPossibility(nums []int) bool {
	n := len(nums)
	counter := 0
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			counter++
			withinBounds := i > 0 && i+2 < n
			if counter > 1 || (withinBounds && nums[i] > nums[i+2] && nums[i-1] > nums[i+1]) {
				return false
			}
		}
	}

	return true
}

// passed, but too long and changes nums.
func checkPossibility2(nums []int) bool {
	counter := 0

	n := len(nums)
	if n == 1 {
		return true
	}
	if nums[1] < nums[0] {
		nums[0] = nums[1]
		counter++
	}

	for i := 2; i < n; i++ {
		if nums[i] < nums[i-1] {
			if i < n-1 {
				if nums[i-2] > nums[i] {
					nums[i] = nums[i-1]
				} else {
					nums[i-1] = nums[i]
				}
			}

			if counter == 1 {
				return false
			}
			counter++
		}
	}

	return counter <= 1
}
