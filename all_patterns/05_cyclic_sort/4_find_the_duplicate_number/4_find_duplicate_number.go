package __find_the_duplicate_number

// We are given an unsorted array containing ‘n+1’ numbers taken from the range 1 to ‘n’.
// The array has only one duplicate but it can be repeated multiple times. Find that duplicate number without using any extra space.
// You are, however, allowed to modify the input array.
func findNumber(nums []int) int {
	for i := 0; i < len(nums); {
		j := nums[i] - 1
		if j == i {
			i++
			continue
		}

		if nums[i] == nums[j] {
			return nums[i]
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	return 0
}

// Solution in O(1) space without modifying input using two pointers
func findNumberTwoPointers(arr []int) int {
	slow, fast := 0, 0
	for {
		slow = arr[slow]
		fast = arr[arr[fast]]

		if slow == fast {
			break
		}
	}

	// find cycle length
	current := arr[slow]
	cycleLen := 0
	for {
		current = arr[current]
		cycleLen++

		if current == arr[slow] {
			break
		}
	}

	return findStart(arr, cycleLen)
}

func findStart(arr []int, cycleLen int) int {
	pointer1, pointer2 := arr[0], arr[0]
	// move pointer2 ahead 'cycleLength' steps
	for cycleLen > 0 {
		pointer2 = arr[pointer2]
		cycleLen--
	}

	// increment both pointers until they meet at the start of the cycle
	for pointer1 != pointer2 {
		pointer1 = arr[pointer1]
		pointer2 = arr[pointer2]
	}

	return pointer1
}
