package _512_number_of_good_pairs

// tptl
// Passed. Should remember this formula: v*(v-1)/2
func numIdenticalPairs(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	res := 0
	for _, v := range m {
		res += v * (v - 1) / 2
	}

	return res
}
