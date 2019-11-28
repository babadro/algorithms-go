package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {

}

// naive
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var res *TreeNode
	recurseTree(root, res, p.Val, q.Val)
	return res
}

func recurseTree(current, res *TreeNode, p, q int) bool {
	if current == nil {
		return false
	}

	var left, right, mid int
	if recurseTree(current.Left, res, p, q) {
		left = 1
	}
	if recurseTree(current.Right, res, p, q) {
		right = 1
	}
	if current.Val == p || current.Val == q {
		mid = 1
	}

	if (left + right + mid) >= 2 {
		res = current
	}

	return (left + right + mid) > 1
}
