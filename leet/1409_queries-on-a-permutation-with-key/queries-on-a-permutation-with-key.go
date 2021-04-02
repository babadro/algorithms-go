package _1409_queries_on_a_permutation_with_key

// passed. todo 1 find more effective solution like
// https://leetcode.com/problems/queries-on-a-permutation-with-key/discuss/972599/C%2B%2B-LRU-implementation
// or
// https://leetcode.com/problems/queries-on-a-permutation-with-key/discuss/601523/Go%3A-Solution-without-rewriting-of-array-on-each-step
func processQueries(queries []int, m int) []int {
	n := len(queries)
	perm := make([]int, m)
	for i := range perm {
		perm[i] = i + 1
	}

	res := make([]int, n)

	var idx int
	for i, num := range queries {
		for j, item := range perm {
			if item == num {
				idx = j
				break
			}
		}

		res[i] = idx

		copy(perm[1:idx+1], perm[0:idx])
		perm[0] = num
	}

	return res
}
