package __binary_tree_path_sum

import "github.com/babadro/algorithms-go/base/binaryTree"

// leetcode 112
func hasPath(root *binaryTree.Node, sum int) bool {
	if root == nil {
		return false
	}

	sum -= root.Val

	if root.Left == nil && root.Right == nil && sum == 0 {
		return true
	}

	return hasPath(root.Left, sum) || hasPath(root.Right, sum)
}
