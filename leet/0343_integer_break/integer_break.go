package _343_integer_break

// passed. tptl. best solution
// https://leetcode.com/problems/integer-break/discuss/80689/A-simple-explanation-of-the-math-part-and-a-O(n)-solution
func integerBreak2(n int) int {
	switch n {
	case 2:
		return 1
	case 3:
		return 2
	}

	res := 1
	for n > 4 {
		res *= 3
		n -= 3
	}

	return res * n
}

// and realize math solution

// passed. bruteforce
func integerBreak(n int) int {
	max := 0
	for count := 2; count <= n; count++ {
		num := n / count
		if num == 0 {
			break
		}

		remainder := n % count

		arr := make([]int, count)
		for i := range arr {
			arr[i] = num
		}

		for i := 0; remainder > 0; i++ {
			i %= len(arr)
			arr[i]++
			remainder--
		}

		product := 1
		for _, number := range arr {
			product *= number
		}

		if product > max {
			max = product
		}
	}

	return max
}
