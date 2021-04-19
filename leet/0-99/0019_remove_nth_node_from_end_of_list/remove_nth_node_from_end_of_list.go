package _019_remove_nth_node_from_end_of_list

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

// TODO 2 implement One Pass algorithm (in solutions)
func removeNthFromEnd(head *single.ListNode, n int) *single.ListNode {
	length := 0
	node := head
	for node != nil {
		length++
		node = node.Next
	}
	if n == length {
		return head.Next
	}
	idx := 0
	node = head
	for idx < length-n-1 {
		node = node.Next
		idx++
	}
	target := node.Next.Next
	node.Next.Next = nil
	node.Next = target
	return head
}
