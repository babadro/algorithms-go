// https://leetcode.com/problems/minimum-number-of-operations-to-make-string-sorted/discuss/1163050/Python-O(26n)-math-solution-beats-100-explained
// https://leetcode.com/problems/minimum-number-of-operations-to-make-string-sorted/discuss/1163811/Java-Solution-O(26n)
// https://leetcode.com/problems/minimum-number-of-operations-to-make-string-sorted/discuss/1164153/C%2B%2B-24-ms

package _1830_minimum_number_of_operations_to_make_string_sorted

// todo 1
func makeStringSortedBruteForce(s string) int {
	b := []byte(s)
	res, n := 0, len(b)

Loop:
	for {
		for i := n - 1; i > 0; i-- {
			if b[i] < b[i-1] {
				res++
				j := i
				for ; j < n-1 && b[j+1] < b[i-1]; j++ {
				}

				b[i-1], b[j] = b[j], b[i-1]
				for k := 0; k < (n-i)/2; k++ {
					idx1, idx2 := k+i, n-1-k
					b[idx1], b[idx2] = b[idx2], b[idx1]
				}

				continue Loop
			}
		}

		return res
	}
}

func reverse(str string, i int) string {
	b := []byte(str)
	for j := 0; j < (len(b)-i)/2; j++ {
		idx1, idx2 := j+i, len(b)-1-j
		b[idx1], b[idx2] = b[idx2], b[idx1]
	}

	return string(b)
}
