package _201_Bitwise_And_of_Numbers_Range

// TODO need to understand
// https://leetcode.com/problems/bitwise-and-of-numbers-range/discuss/56808/Share-my-C%2B%2B-solution-with-explanationeasy-to-understand
func rangeBitwiseAnd(m int, n int) int {
	i := 0
	for m != n {
		m = m >> 1
		n = n >> 1
		i++
	}
	return m << i
}

// naive
func rangeBitwiseAndNaive(m int, n int) int {
	res := m
	for i := m + 1; i <= n; i++ {
		res = res & i
	}
	return res
}
