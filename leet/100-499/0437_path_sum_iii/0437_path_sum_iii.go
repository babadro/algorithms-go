package _437_path_sum_iii

import "github.com/babadro/algorithms-go/base/binaryTree"

// dyx passed, but too slow. todo 2 find faster solution
func pathSum(root *binaryTree.Node, targetSum int) int {
	res := 0
	f := func(n *binaryTree.Node) {
		traverse(n, 0, targetSum, &res)
	}

	inOrder(root, f)

	return res
}

func traverse(node *binaryTree.Node, cur, target int, res *int) {
	if node == nil {
		return
	}

	cur += node.Val
	if cur == target {
		*res++
	}

	traverse(node.Left, cur, target, res)
	traverse(node.Right, cur, target, res)
}

func inOrder(node *binaryTree.Node, f func(n *binaryTree.Node)) {
	if node != nil {
		inOrder(node.Left, f)
		f(node)
		inOrder(node.Right, f)
	}
}
