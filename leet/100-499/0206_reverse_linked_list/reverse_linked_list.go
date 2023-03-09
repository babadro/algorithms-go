package _206_reverse_linked_list

import (
	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/singly"
)

func reverseList(head *singly.ListNode) *singly.ListNode {
	var prev *singly.ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

func reverseListRecursive(head *singly.ListNode) *singly.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return tmp
}
