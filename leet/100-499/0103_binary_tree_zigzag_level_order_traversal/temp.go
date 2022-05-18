package _103_binary_tree_zigzag_level_order_traversal

import "github.com/babadro/algorithms-go/base/binaryTree"

func zigzagLevelOrder3(root *binaryTree.Node) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	q := []*binaryTree.Node{root}
	for len(q) > 0 {
		level := make([]int, len(q))
		for i := 0; i < len(q); i++ {
			node := q[0]
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

			level = append(level, node.Val)
			q = q[1:]
		}

		res = append(res, level)
	}

	return res
}
