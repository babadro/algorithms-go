package _876_middle_of_the_linked_list

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

// tptl. passed. fast and slow pointers
func middleNode2(head *single.ListNode) *single.ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func middleNode(head *single.ListNode) *single.ListNode {
	l := 0
	for node := head; node != nil; node = node.Next {
		l++
	}

	res := head
	for i := 0; i < l/2; i++ {
		res = res.Next
	}

	return res
}
