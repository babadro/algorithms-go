package _1837_sum_of_digits_in_base_k

// tptl. passed. medium for me.
// https://leetcode.com/problems/sum-of-digits-in-base-k/discuss/1179863/Go-golang-solution
func sumBase(n int, k int) int {
	res := 0
	for n > 0 {
		res += n % k
		n /= k
	}

	return res
}
