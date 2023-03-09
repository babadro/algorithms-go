package _203_remove_linked_list_elements

import (
	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/singly"
)

// tptl. best solution
func removeElements2(head *singly.ListNode, val int) *singly.ListNode {
	dummy := &singly.ListNode{}
	head, dummy.Next = dummy, head
	for dummy.Next != nil {
		if dummy.Next.Val == val {
			dummy.Next = dummy.Next.Next
		} else {
			dummy = dummy.Next
		}
	}

	return head.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// passed, but it has too many lines of code.
func removeElements(head *singly.ListNode, val int) *singly.ListNode {
	node := head
	for node != nil && node.Val == val {
		node = node.Next
	}

	if node == nil {
		return nil
	}

	head = node

	prev, next := head, head.Next
	for next != nil {
		if next.Val == val {
			prev.Next = next.Next
			next = next.Next
		} else {
			prev, next = next, next.Next
		}
	}

	return head
}
