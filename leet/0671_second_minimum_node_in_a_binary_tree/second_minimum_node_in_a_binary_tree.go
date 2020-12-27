package _671_second_minimum_node_in_a_binary_tree

import (
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"
	"math"
	"sort"
)

// passed. best solution. tptl
func findSecondMinimumValue2(root *btree.Node) int {
	res := visit2(root, root.Val)

	if res == math.MaxInt64 {
		return -1
	}

	return res
}

func visit2(root *btree.Node, val int) int {
	if root == nil {
		return math.MaxInt64
	}

	if root.Val > val {
		return root.Val
	}

	left, right := visit2(root.Left, val), visit2(root.Right, val)

	if left < right {
		return left
	}

	return right
}

// passed bruteforce
func findSecondMinimumValue(root *btree.Node) int {
	m := make(map[int]bool)
	visit(root, m)
	arr := make([]int, 0, len(m))
	for num := range m {
		arr = append(arr, num)
	}
	if len(arr) < 2 {
		return -1
	}

	sort.Ints(arr)
	return arr[1]
}

func visit(root *btree.Node, m map[int]bool) {
	if root != nil {
		visit(root.Left, m)
		m[root.Val] = true
		visit(root.Right, m)
	}
}
