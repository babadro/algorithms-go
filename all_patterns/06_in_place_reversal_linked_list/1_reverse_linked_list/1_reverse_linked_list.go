package __reverse_linked_list

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

// leetcode 206
func reverseLinkedList(node *single.ListNode) *single.ListNode {
	var prev, next *single.ListNode
	for node != nil {
		next = node.Next
		node.Next = prev

		prev, node = node, next
	}

	return prev
}
