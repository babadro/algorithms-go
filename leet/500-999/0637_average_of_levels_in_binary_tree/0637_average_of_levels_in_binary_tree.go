package _637_average_of_levels_in_binary_tree

import "github.com/babadro/algorithms-go/base/binaryTree"

// dyx. passed
func averageOfLevels(root *binaryTree.Node) []float64 {
	if root == nil {
		return nil
	}

	q := []*binaryTree.Node{root}
	var res []float64
	for len(q) > 0 {
		levelSize := len(q)
		sum := 0
		for i := 0; i < levelSize; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

			sum += node.Val
		}

		res = append(res, float64(sum)/float64(levelSize))
	}

	return res
}
