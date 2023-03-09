package __reverse_linked_list

import (
	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/singly"
)

// leetcode 206
func reverseLinkedList(node *singly.ListNode) *singly.ListNode {
	var prev, next *singly.ListNode
	for node != nil {
		next = node.Next
		node.Next = prev

		prev, node = node, next
	}

	return prev
}
