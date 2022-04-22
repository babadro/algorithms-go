package _1_linked_list_cycle

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

// leetcode 141
func hasCycle(node *single.ListNode) bool {
	fast, slow := node, node
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}

	return false
}

// tptl
func findCycleLen(node *single.ListNode) int {
	fast, slow := node, node
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow { // has cycle
			cur, length := slow, 0
			for {
				cur = cur.Next
				length++

				if cur == slow {
					break
				}
			}

			return length
		}
	}

	return 0
}
