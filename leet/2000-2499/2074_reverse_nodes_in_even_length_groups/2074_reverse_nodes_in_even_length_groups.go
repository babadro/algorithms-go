package _2074_reverse_nodes_in_even_length_groups

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

func reverseEvenLengthGroups(head *single.ListNode) *single.ListNode {
	if head == nil {
		return nil
	}

	cur, prev := head, head
	for groupSize := 1; cur != nil; groupSize++ {
		if groupSize%2 == 1 {
			for i := 0; i < groupSize && cur != nil; i++ {
				prev = cur
				cur = cur.Next
			}

			continue
		}

		left := prev
		var right *single.ListNode
		for i := 0; i < groupSize && cur != nil; i++ {
			right = cur.Next
			cur.Next = prev
			cur = right
		}
	}
}
