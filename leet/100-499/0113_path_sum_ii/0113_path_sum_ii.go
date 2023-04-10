package _113_path_sum_ii

import "github.com/babadro/algorithms-go/base/binaryTree"

// dyx. passed.
// todo 2 iterative solution
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

// version with optimized cur array for reusing in recursive solution
func pathSum2(root *binaryTree.Node, targetSum int) [][]int {
	var res [][]int

	rec(root, &[]int{}, &res, 0, targetSum)

	return res
}

func rec(node *binaryTree.Node, arr *[]int, res *[][]int, arrLen, target int) {
	if node == nil {
		return
	}

	target -= node.Val
	*arr = append((*arr)[:arrLen], node.Val)
	arrLen++

	if node.Left == nil && node.Right == nil {
		if target == 0 {
			arrCopy := make([]int, len(*arr))
			copy(arrCopy, *arr)
			*res = append(*res, arrCopy)
		}

		return
	}

	rec(node.Left, arr, res, arrLen, target)
	rec(node.Right, arr, res, arrLen, target)
}
