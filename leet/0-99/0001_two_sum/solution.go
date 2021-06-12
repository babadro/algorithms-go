package _001_two_sum

// one-pass hash table. tptl. easy (hard for me). #array
func twoSum3(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for idx1, num1 := range nums {
		num2 := target - num1
		if idx2, ok := m[num2]; ok {
			return []int{idx1, idx2}
		}

		m[num1] = idx1
	}

	return []int{}
}

// passed. two-pass hash table
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		m[num] = i
	}

	for i, num := range nums {
		if idx, ok := m[target-num]; ok && idx != i {
			return []int{i, idx}
		}
	}

	return nil
}
