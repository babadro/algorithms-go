package _013_partition_array_into_three_parts_with_equal_sum

// tptl. best solution
func canThreePartsEqualSum2(arr []int) bool {
	total := 0
	for _, num := range arr {
		total += num
	}
	if total%3 != 0 {
		return false
	}

	temp, count := 0, 0
	for i := 0; i < len(arr)-1; i++ { // i < len(arr)-1 not i < len(arr), important!
		temp += arr[i]
		if temp == total/3 {
			temp = 0
			count++
		}

		if count == 2 {
			return true
		}
	}

	return false
}

// passed
func canThreePartsEqualSum(arr []int) bool {
	total := 0
	for _, num := range arr {
		total += num
	}
	if total%3 != 0 {
		return false
	}

	partCount := total / 3
	sum := 0
	idx1 := -1
	for i := range arr {
		sum += arr[i]
		if sum == partCount {
			idx1 = i
			break
		}
	}

	if idx1 == -1 {
		return false
	}

	sum = 0
	for j := len(arr) - 1; j > idx1+1; j-- {
		sum += arr[j]
		if sum == partCount {
			return true
		}
	}

	return false
}
