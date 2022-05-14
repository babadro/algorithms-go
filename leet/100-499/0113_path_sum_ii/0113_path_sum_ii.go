package _113_path_sum_ii

import "github.com/babadro/algorithms-go/base/binaryTree"

// dyx. passed.
// todo 2 iterative solution
// todo 3 optimize cur array for reusing in recursive solution
func pathSum(root *binaryTree.Node, targetSum int) [][]int {
	var res [][]int
	preOrder(root, 0, targetSum, nil, &res)

	return res
}

func preOrder(node *binaryTree.Node, sum, target int, cur []int, res *[][]int) {
	if node != nil {
		cur = append(cur, node.Val)
		sum += node.Val

		if node.Left == nil && node.Right == nil {
			if sum == target {
				newPath := make([]int, len(cur))
				copy(newPath, cur)
				*res = append(*res, newPath)
			}
		} else {
			preOrder(node.Left, sum, target, cur, res)
			preOrder(node.Right, sum, target, cur, res)
		}
	}
}
