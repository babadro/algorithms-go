package _124_binary_tree_maximum_path_sum

import (
	"math"

	"github.com/babadro/algorithms-go/base/binaryTree"
)

// dyx passed
func maxPathSum(root *binaryTree.Node) int {
	res := math.MinInt64
	rec(root, &res)

	return res
}

func rec(node *binaryTree.Node, res *int) int {
	if node == nil {
		return 0
	}

	leftRes := max(0, rec(node.Left, res))
	rightRes := max(0, rec(node.Right, res))
	path := leftRes + rightRes + node.Val
	*res = max(*res, path)

	return node.Val + max(leftRes, rightRes)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
