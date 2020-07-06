package _116_populating_next_right_pointers_in_each_node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

var Node1 *Node
var Node2 *Node
var Counter = 0
var Sequence = 1

// TODO 1 cannot allocate stack error
func connect(root *Node) *Node {
	level := 1

	for traversalLevel(root, level) {
		level++
	}
	return root
}

// TODO 2 need to understand
func traversalLevel(node *Node, level int) bool {
	if node == nil {
		return false
	}

	if level == 1 {
		Node2 = Node1
		Node1 = node
		Counter++
		if Counter != Sequence {
			Node1.Next = Node2
		} else {
			Sequence *= 2
		}
		return true
	}

	right := traversalLevel(node.Right, level-1)
	left := traversalLevel(node.Left, level-1)

	return left || right
}
