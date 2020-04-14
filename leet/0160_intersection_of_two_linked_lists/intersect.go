package _160_intersection_of_two_linked_lists

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

func getIntersectionNode(headA, headB *single.ListNode) *single.ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	node1 := headA
	node2 := headB
	var lastA, lastB *single.ListNode
	for node1 != node2 {
		if lastA != nil && lastB != nil && lastA != lastB {
			return nil
		}
		if node1.Next == nil {
			lastA = node1
			node1 = headB
		} else {
			node1 = node1.Next
		}
		if node2.Next == nil {
			lastB = node2
			node2 = headA
		} else {
			node2 = node2.Next
		}
	}
	return node1
}
