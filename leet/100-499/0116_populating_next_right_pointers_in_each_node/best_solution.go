package _116_populating_next_right_pointers_in_each_node

func connectBest(root *Node) *Node {
	connectNext(root)
	return root
}

func connectNext(root *Node) {
	if root == nil || root.Left == nil {
		return
	}

	root.Left.Next = root.Right

	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}
	connectNext(root.Left)
	connectNext(root.Right)
}
