package _061_rotate_list

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

func rotateRight3(head *single.ListNode, k int) *single.ListNode {
	listLen := 0
	tail := head
	for node := head; node != nil; listLen++ {
		tail = node
		node = node.Next
	}

	if listLen == 0 {
		return nil
	}

	k %= listLen
	if k == 0 {
		return head
	}

	newTail := head
	for i := 1; i < listLen-k; i++ {
		newTail = newTail.Next
	}

	newHead := newTail.Next
	newTail.Next = nil
	tail.Next = head

	return newHead
}
