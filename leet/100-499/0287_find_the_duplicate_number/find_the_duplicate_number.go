package _287_find_the_duplicate_number

// 31% 13% - it's ok for two floyd cycle solution
// https://medium.com/@simrangarg0501/finding-the-duplicate-number-using-floyds-tortoise-and-hare-algorithm-618ced80e98e
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return fast
}

// "classical" fast and slow pointer solution
func findDuplicate2(nums []int) int {
	fast, slow := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if fast == slow {
			break
		}
	}

	cycleLen := 0
	for {
		slow = nums[slow]
		cycleLen++

		if slow == fast {
			break
		}
	}

	i1 := 0
	for i := 0; i < cycleLen; i++ {
		i1 = nums[i1]
	}

	for i2 := 0; i2 != i1; {
		i1 = nums[i1]
		i2 = nums[i2]
	}

	return i1
}

// cyclic sort (with modifying input array)
func findDuplicate3(nums []int) int {
	for i := range nums {
		for {
			idx := nums[i] - 1
			if nums[i] == nums[idx] {
				break
			}

			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}

	for i, num := range nums {
		if num != i+1 {
			return num
		}
	}

	return 0
}
