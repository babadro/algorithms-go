package _2213_longest_subsring_of_one_repeating_character

// https://leetcode.com/problems/longest-substring-of-one-repeating-character/discuss/1877448/Merge-intervals-using-map
func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	m, sizes := make(map[int]int), make(map[int]int)
	res := make([]int, 0, len(queryIndices))
}

type maxHeap []int

func (m maxHeap) Len() int           { return len(m) }
func (m maxHeap) Less(i, j int) bool { return m[i] > m[j] }
func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}
func (m *maxHeap) Pop() interface{} {
	last := len(*m) - 1
	res := (*m)[last]
	*m = (*m)[:last]

	return res
}
