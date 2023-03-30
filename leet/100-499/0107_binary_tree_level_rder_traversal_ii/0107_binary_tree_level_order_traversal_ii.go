package _107_binary_tree_level_rder_traversal_ii

import "github.com/babadro/algorithms-go/base/binaryTree"

func levelOrderBottom(root *binaryTree.Node) [][]int {
	if root == nil {
		return nil
	}

	q := []*binaryTree.Node{root}
	var res [][]int
	for len(q) > 0 {
		levelLen := len(q)
		level := make([]int, 0, levelLen)
		for ; levelLen > 0; levelLen-- {
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

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
