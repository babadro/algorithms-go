package btree

import "container/list"

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

// dyx
func LevelOrderFuncRecursive(node *Node, f func(node *Node)) {
	h := Height(node)

	for i := 1; i <= h; i++ {
		levelOrderFuncHelper(node, i, f)
	}
}

func levelOrderFuncHelper(node *Node, level int, f func(node *Node)) {
	if node == nil {
		return
	}

	if level == 1 {
		f(node)
	} else if level > 1 {
		levelOrderFuncHelper(node.Left, level-1, f)
		levelOrderFuncHelper(node.Right, level-1, f)
	}
}

// dyx
func LevelOrderFuncIterative(node *Node, f func(node *Node)) {
	if node == nil {
		return
	}

	q := list.New()
	q.PushBack(node)

	for q.Front() != nil {
		tempNode := q.Front().Value.(*Node)
		q.Remove(q.Front())

		f(tempNode)

		if tempNode.Left != nil {
			q.PushBack(tempNode.Left)
		}

		if tempNode.Right != nil {
			q.PushBack(tempNode.Right)
		}
	}
}

// dyx
func Height(node *Node) int {
	if node == nil {
		return 0
	}

	l, r := Height(node.Left), Height(node.Right)
	if l > r {
		return l + 1
	}

	return r + 1
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
