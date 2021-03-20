package _1018_binary_prefix_divisible_by_5

// tptl. passed. not easy
func prefixesDivBy5(A []int) []bool {
	num := 0
	res := make([]bool, len(A))
	for i, item := range A {
		if num = (2*num + item) % 5; num == 0 {
			res[i] = true
		}
	}

	return res
}
