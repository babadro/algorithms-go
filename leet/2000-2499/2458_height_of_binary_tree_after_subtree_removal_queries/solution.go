package _2458_height_of_binary_tree_after_subtree_removal_queries

import "github.com/babadro/algorithms-go/base/binaryTree"

var seen [100002]int
var curMax int

// best solution
func treeQueries2(root *binaryTree.Node, queries []int) []int {
	seen = [100002]int{} // the height seen so far when visiting ith node

	curMax = 0
	dfs(root, 0)

	curMax = 0
	dfs(root, 0)

	r := make([]int, len(queries))
	for i, q := range queries {
		r[i] = seen[q]
	}

	return r
}

func dfs(p *binaryTree.Node, h int) {
	if seen[p.Val] < curMax {
		seen[p.Val] = curMax
	}

	if curMax < h {
		curMax = h
	}

	if p.Left != nil {
		dfs(p.Left, h+1)
	}

	if p.Right != nil {
		dfs(p.Right, h+1)
	}

	p.Left, p.Right = p.Right, p.Left
}
