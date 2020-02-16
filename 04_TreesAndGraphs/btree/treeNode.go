package btree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func InOrder(node *Node, list *[]int) {
	if node != nil {
		InOrder(node.Left, list)
		*list = append(*list, node.Val)
		InOrder(node.Right, list)
	}
}

func InOrderFunc(node *Node, f func(node *Node)) {
	if node != nil {
		InOrderFunc(node.Left, f)
		f(node)
		InOrderFunc(node.Right, f)
	}
}
