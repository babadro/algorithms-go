package _543_tree_diameter

import "github.com/babadro/algorithms-go/base/binaryTree"

// dyx passed
func diameterOfBinaryTree(root *binaryTree.Node) int {
	var res int
	rec(root, &res)

	return res
}

func rec(node *binaryTree.Node, res *int) int {
	if node == nil {
		return 0
	}

	leftRes, rightRes := rec(node.Left, res), rec(node.Right, res)
	diameter := leftRes + rightRes
	*res = max(*res, diameter)

	return 1 + max(leftRes, rightRes)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
