package _645_set_mismatch

// passed. best solution
func findErrorNums3(nums []int) []int {
	for i := 0; i < len(nums); {
		idx := nums[i] - 1
		if nums[i] == nums[idx] {
			i++
			continue
		}

		nums[i], nums[idx] = nums[idx], nums[i]
	}

	for i, num := range nums {
		if num != i+1 {
			return []int{num, i + 1}
		}
	}

	return nil
}

// passed. tptl. todo 2 look for xor solution
func findErrorNums(nums []int) []int {
	n := len(nums)
	sumArithmeticProgression := (1 + n) * n / 2

	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	diff := sumArithmeticProgression - sum

	for i := 0; i < n; {
		idx := nums[i] - 1

		if idx == i {
			i++
			continue
		}

		if nums[i] == nums[idx] {
			return []int{nums[i], nums[i] + diff}
		}

		nums[i], nums[idx] = nums[idx], nums[i]
	}

	return nil
}

// doesn't work
func findErrorNums2(nums []int) []int {
	i, k, j, duplicate := 0, 0, 0, len(nums)
	for i < duplicate {
		idx := nums[i] - 1
		if idx == i {
			i++
			continue
		}

		if nums[i] == nums[idx] {
			duplicate = nums[i]
			k, j = i, idx
			i++
			continue
		}

		nums[i], nums[idx] = nums[idx], nums[i]
	}

	if j == nums[k] {
		return []int{k + 1, j + 1}
	}

	return []int{j + 1, k + 1}
}
