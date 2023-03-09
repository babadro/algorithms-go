package _138_copy_list_with_random_pointer

import (
	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/singly"
)

// 100% 80%
func copyRandomList(head *singly.ListNodeRandom) *singly.ListNodeRandom {
	if head == nil {
		return nil
	}
	var copyNodes []*singly.ListNodeRandom

	node := head
	idx := 0
	for node != nil {
		copyNodes = append(copyNodes, &singly.ListNodeRandom{Val: node.Val})
		node.Val = idx
		idx++
		node = node.Next
	}

	node = head
	idx = 0
	for node != nil {
		if node.Random != nil {
			randomIdx := node.Random.Val
			copyNodes[idx].Random = copyNodes[randomIdx]
		}
		if node.Next != nil {
			copyNodes[idx].Next = copyNodes[idx+1]
		}
		idx++
		node = node.Next
	}

	return copyNodes[0]
}
