package _2183_count_array_pairs_divisible_by_k

// tptl. hard. passed. not mine.
// todo 2 look at the Approach 2: Optimization here:
// https://leetcode.com/problems/count-array-pairs-divisible-by-k/discuss/1784721/Count-GCDs
func countPairs(nums []int, k int) int64 {
	res := 0
	gcds := make(map[int]int)
	for _, num := range nums {
		gcdI := gcd(num, k)
		for gcdJ, cnt := range gcds {
			if gcdI*gcdJ%k == 0 {
				res += cnt
			}
		}

		gcds[gcdI]++
	}

	return int64(res)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}
