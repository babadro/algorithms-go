package _230_kth_smallest_element_in_bst

import "github.com/babadro/algorithms-go/04_TreesAndGraphs/binaryTree"

// 54% 51%
func kthSmallest(root *binaryTree.Node, k int) int {
	res, counter := 0, 0
	inOrder(root, &counter, &res, k)
	return res
}

func inOrder(root *binaryTree.Node, counter, res *int, k int) {
	if root != nil {
		inOrder(root.Left, counter, res, k)
		if *counter >= k {
			return
		}
		*res = root.Val
		*counter++
		inOrder(root.Right, counter, res, k)
	}
}

// todo 2 doesn't work. fix it. look at solution.
func kthSmallestIterative(root *binaryTree.Node, k int) int {
	var stack []*binaryTree.Node

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if k == 0 {
			return root.Val
		}
		k--
		root = root.Right
	}
}
