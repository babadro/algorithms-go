package _141_linked_list_cycle

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *single.ListNode) bool {
	nodeMap := make(map[*single.ListNode]bool)
	node := head
	for node != nil {
		if ok := nodeMap[node]; ok {
			return true
		}
		nodeMap[node] = true
		node = node.Next
	}
	return false
}

// tptl. passed
func hasCycleTwoPointers(head *single.ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}

	return false
}
