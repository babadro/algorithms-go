package _092_reverse_linked_list_ii

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

// tptl. passed. medium #linkedlist
func reverseBetween2(head *single.ListNode, left int, right int) *single.ListNode {
	leftNode := head
	var prevLeft *single.ListNode
	for i := 1; i < left; i++ {
		prevLeft, leftNode = leftNode, leftNode.Next
	}

	// reverse
	var prev *single.ListNode
	cur := leftNode
	for i := left; i <= right; i++ {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}

	rightNode, nextRight := prev, cur
	leftNode, rightNode = rightNode, leftNode

	if prevLeft != nil {
		prevLeft.Next = leftNode
	} else {
		head = leftNode
	}

	rightNode.Next = nextRight

	return head
}
