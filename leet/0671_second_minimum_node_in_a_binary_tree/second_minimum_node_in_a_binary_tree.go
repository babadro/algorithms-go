package _671_second_minimum_node_in_a_binary_tree

import (
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"
	"sort"
)

// passed bruteforce
// todo 1 make more effective solution
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
