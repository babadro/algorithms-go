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

func PreOrderFunc(node *Node, f func(node *Node)) {
	if node != nil {
		f(node)
		PreOrderFunc(node.Left, f)
		PreOrderFunc(node.Right, f)
	}
}

// tptl
func Clone(root *Node) *Node {
	if root != nil {
		n := &Node{Val: root.Val}
		n.Left = Clone(root.Left)
		n.Right = Clone(root.Right)

		return root
	}

	return nil
}

func Delete(root *Node, key int) *Node {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = Delete(root.Left, key)
	} else if key > root.Val {
		root.Right = Delete(root.Right, key)
	} else {
		if root.Right == nil {
			return root.Left
		} else if root.Left == nil {
			return root.Right
		}

		rightSmallest := root.Right
		for rightSmallest.Left != nil {
			rightSmallest = rightSmallest.Left
		}

		rightSmallest.Left = root.Left

		return root.Right
	}

	return root
}
