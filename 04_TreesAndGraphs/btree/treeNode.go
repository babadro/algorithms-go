package btree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InOrder(node *TreeNode, list *[]int) {
	if node != nil {
		InOrder(node.Left, list)
		*list = append(*list, node.Val)
		InOrder(node.Right, list)
	}
}
