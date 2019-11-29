package _235_lowest_common_ancestor_bst

import (
	bst "algorithms-go/leet/0108_sorted_array_to_BST"
	"fmt"
)

// naive
func LowestCommonAncestor(root, p, q *bst.TreeNode) *bst.TreeNode {
	var res *bst.TreeNode
	recurseTree(root, res, p.Val, q.Val)
	return res
}

func recurseTree(current, res *bst.TreeNode, p, q int) bool {
	if current == nil {
		return false
	}
	fmt.Println(current.Val)

	var left, right, mid int
	if recurseTree(current.Left, res, p, q) {
		left = 1
	}
	if recurseTree(current.Right, res, p, q) {
		right = 1
	}
	if current.Val == p || current.Val == q {
		mid = 1
	}

	if (left + right + mid) >= 2 {
		fmt.Println("bingo")
		res = current
	}

	return (left + right + mid) > 1
}
