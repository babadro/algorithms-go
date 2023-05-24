package _2458_height_of_binary_tree_after_subtree_removal_queries

import "github.com/babadro/algorithms-go/base/binaryTree"

var seen [100002]int
var curMax int

// best solution
func treeQueries2(root *binaryTree.Node, queries []int) []int {
	seen = [100002]int{} // the height seen so far when visiting ith node

	curMax = 0
	dfs(root, 0, false)

	curMax = 0
	dfs(root, 0, true)

	r := make([]int, len(queries))
	for i, q := range queries {
		r[i] = seen[q]
	}

	return r
}

func dfs(p *binaryTree.Node, h int, revert bool) {
	if seen[p.Val] < curMax {
		seen[p.Val] = curMax
	}

	if curMax < h {
		curMax = h
	}

	first, second := p.Left, p.Right
	if revert {
		first, second = p.Right, p.Left
	}

	if first != nil {
		dfs(first, h+1, revert)
	}

	if second != nil {
		dfs(second, h+1, revert)
	}
}
