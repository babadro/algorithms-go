package _102_binary_tree_level_order_traversal

import "github.com/babadro/algorithms-go/base/binaryTree"

// tptl. passed
func levelOrder2(root *binaryTree.Node) [][]int {
	if root == nil {
		return nil
	}

	q := []*binaryTree.Node{root}
	var res [][]int
	for len(q) > 0 {
		level := make([]int, 0, len(q))

		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		res = append(res, level)
	}

	return res
}
