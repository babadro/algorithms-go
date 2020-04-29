package _283_move_zeroes

func moveZeroes(nums []int) {
	zeroesCount, pos := 0, 0

	for _, num := range nums {
		if num == 0 {
			zeroesCount++
		} else {
			nums[pos] = num
			pos++
		}
	}

	for i := len(nums) - 1; i >= 0 && zeroesCount > 0; i-- {
		nums[i] = 0
		zeroesCount--
	}
}

func moveZeroesNaive(nums []int) {
	zeroExists := false
	for i := range nums {
		if nums[i] == 0 {
			zeroExists = true
			continue
		} else if zeroExists {
			for j := i; j > 0; j-- {
				if nums[j-1] != 0 {
					break
				}
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}
