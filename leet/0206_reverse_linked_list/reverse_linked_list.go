package _206_reverse_linked_list

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

func reverseList(head *single.ListNode) *single.ListNode {
	node := head
	var prev *single.ListNode
	var next *single.ListNode
	for node != nil {
		next = node.Next
		node.Next = prev
		prev = node
		node = next
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
