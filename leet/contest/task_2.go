package contest

type node struct {
	val   int
	left  *node
	right *node
}

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	root := makeBTree(leftChild, rightChild)
	nodeMap := make(map[*node]bool)
	counter := 0
	valid := true

	f := func(n *node) bool {
		if _, exists := nodeMap[n]; exists {
			valid = false
			return false
		}
		nodeMap[n] = true
		counter++
		return true
	}

	PreOrderFunc(root, f)
	if counter != n {
		valid = false
	}
	return valid
}

func makeBTree(leftChild, rightChild []int) *node {
	nodes := make([]*node, len(leftChild))

	for i := range nodes {
		nodes[i] = &node{val: i}
	}

	for i := range leftChild {
		if childIdx := leftChild[i]; childIdx != -1 {
			nodes[i].left = nodes[childIdx]
		}
	}

	for i := range rightChild {
		if childIdx := rightChild[i]; childIdx != -1 {
			nodes[i].right = nodes[childIdx]
		}
	}

	return nodes[0]
}

func PreOrderFunc(node *node, f func(node *node) bool) {
	if node != nil {
		if !f(node) {
			return
		}
		PreOrderFunc(node.left, f)
		PreOrderFunc(node.right, f)
	}
}
