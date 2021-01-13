package _112_path_sum

import "github.com/babadro/algorithms-go/base/binaryTree"

// best solution. tptl
func hasPathSum2(root *binaryTree.Node, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}

	return hasPathSum2(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
