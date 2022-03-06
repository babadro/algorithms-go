package _2176_count_equal_and_divisible_pairs_in_an_array

// tptl passed
// https://leetcode.com/problems/count-equal-and-divisible-pairs-in-an-array/discuss/1784521/O(n-sqrt-n)
func countPairs(nums []int, k int) int {
	res := 0
	m := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = append(m[nums[i]], i)
	}

	for _, ids := range m {
		gcds := make(map[int]int)
		for _, i := range ids {
			gcdI := gcd(i, k)
			for gcdJ, cnt := range gcds {
				if gcdI*gcdJ%k == 0 {
					res += cnt
				}
			}

			gcds[gcdI]++
		}
	}

	return res
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}
