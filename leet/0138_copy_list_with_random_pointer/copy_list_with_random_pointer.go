package _138_copy_list_with_random_pointer

type arrElem struct {
	node *Node
	idx int
	randIdx int
}

type Node struct {
	    Val int
	    Next *Node
	    Random *Node
	}

// TODO 1
// maybe use node.val as randIdx...
func copyRandomList(head *Node) *Node {
	var arr []arrElem

	node := head
	idx := 0
	for node != nil {
		arr = append(arr, arrElem{node: node, idx: idx})
		idx++
	}
	for i := range arr {
		random := arr[i].node.Random
		if random != nil {
			arr[i].randIdx =
		}
	}
}
