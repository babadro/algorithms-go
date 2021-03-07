package _785_minimum_elements_to_add_to_form_a_given_sum

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	diff := sum - goal
	if sum < goal {
		diff = goal - sum
	}

	if diff%limit != 0 {
		return diff/limit + 1
	}

	return diff / limit
}
