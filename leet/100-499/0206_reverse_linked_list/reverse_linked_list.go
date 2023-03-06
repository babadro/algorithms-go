package _206_reverse_linked_list

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

func reverseList(head *single.ListNode) *single.ListNode {
	var prev *single.ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

func reverseListRecursive(head *single.ListNode) *single.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return tmp
}
