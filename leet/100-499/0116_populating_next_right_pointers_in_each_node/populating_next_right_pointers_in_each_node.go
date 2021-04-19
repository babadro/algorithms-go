package _116_populating_next_right_pointers_in_each_node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 94% 26%
func connect(root *Node) *Node {
	level := 1
	nodePair := []*Node{&Node{}, &Node{}}
	counter, sequence := 0, 1
	for traversalLevel(root, nodePair, level, &counter, &sequence) {
		level++
	}
	return root
}

// TODO 2 need to understand
func traversalLevel(node *Node, nodePair []*Node, level int, counter, sequence *int) bool {
	if node == nil {
		return false
	}

	if level == 1 {
		nodePair[1] = nodePair[0]
		nodePair[0] = node
		*counter++
		if *counter != *sequence {
			nodePair[0].Next = nodePair[1]
		} else {
			*sequence *= 2
		}
		return true
	}

	right := traversalLevel(node.Right, nodePair, level-1, counter, sequence)
	left := traversalLevel(node.Left, nodePair, level-1, counter, sequence)

	return left || right
}
