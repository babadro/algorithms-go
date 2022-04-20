package _1_pair_with_target_sum

func search(sortedArr []int, target int) []int {
	for i, j := 0, len(sortedArr)-1; i < j; {
		if sum := sortedArr[i] + sortedArr[j]; sum == target {
			return []int{i, j}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}

	return nil
}

// alternative approach (not two pointers)
func searchHashMap(sortedArr []int, target int) []int {
	nums := make(map[int]int)
	for i := range sortedArr {
		if idx, ok := nums[target-sortedArr[i]]; ok {
			return []int{idx, i}
		} else {
			nums[sortedArr[i]] = i
		}
	}

	return nil
}
