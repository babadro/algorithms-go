package _1803_count_pairs_with_xor_in_a_range

import "github.com/babadro/algorithms-go/utils"

// because constraint: 1 <= nums[i] <= 2 * 10^4
const bitCount = 14

type trieNode struct {
	cnt   int
	links [2]*trieNode
}

func (tn *trieNode) insert(n, i int) {
	node := tn
	for i >= 0 {
		b := (n >> i) & 1
		if node.links[b] == nil {
			node.links[b] = &trieNode{}
		}

		node.links[b].cnt++
		i--
		node = node.links[b]
	}
}

func (tn *trieNode) countLess(n, lim, i int) int {
	b := (n >> utils.Abs(i)) & 1
	limB := (lim >> utils.Abs(i)) & 1
	x := b ^ limB

	res := 0
	if tn.links[b] != nil {
		res = limB * tn.links[b].cnt
	}

	if tn.links[x] != nil {
		res += tn.links[x].countLess(n, lim, i-1)
	}

	return res
}

// todo 1. passed. tptl. hard. need to understand
// https://leetcode.com/problems/count-pairs-with-xor-in-a-range/discuss/1122495/C%2B%2B-with-picture
func countPairs(nums []int, low int, high int) int {
	var root trieNode
	res := 0
	for _, num := range nums {
		res += root.countLess(num, high+1, bitCount) - root.countLess(num, low, bitCount)
		root.insert(num, bitCount)
	}

	return res
}

// todo 2
// non trie solution: https://leetcode.com/problems/count-pairs-with-xor-in-a-range/discuss/1124877/picture-explanation-for-the-10-line-O(n)-solution-python3
