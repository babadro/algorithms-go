package _129_sum_root_to_leaf_numbers

import "github.com/babadro/algorithms-go/base/binaryTree"

// passed. dyx.
// todo 2 return value from preOrder instead of using pointer
func sumNumbers(root *binaryTree.Node) int {
	sum := 0
	preOrder(root, 0, &sum)

	return sum
}

func preOrder(node *binaryTree.Node, cur int, sum *int) {
	if node != nil {
		cur *= 10
		cur += node.Val

		if node.Left == nil && node.Right == nil {
			*sum += cur
		} else {
			preOrder(node.Left, cur, sum)
			preOrder(node.Right, cur, sum)
		}
	}
}
