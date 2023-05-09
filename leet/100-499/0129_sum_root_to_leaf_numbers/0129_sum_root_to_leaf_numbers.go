package _129_sum_root_to_leaf_numbers

import "github.com/babadro/algorithms-go/base/binaryTree"

// passed. dyx.
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

func sumNumbers2(root *binaryTree.Node) int {
	return preOrder2(root, 0)
}

func preOrder2(node *binaryTree.Node, cur int) int {
	if node == nil {
		return 0
	}

	cur = cur*10 + node.Val

	if node.Left == nil && node.Right == nil {
		return cur
	}

	return preOrder2(node.Left, cur) + preOrder2(node.Right, cur)
}
