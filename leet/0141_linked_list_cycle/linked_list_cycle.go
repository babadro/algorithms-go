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
