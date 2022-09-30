package _2074_reverse_nodes_in_even_length_groups

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

// todo 1 doesn't work
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

		left, start := prev, cur
		for i := 0; i < groupSize && cur != nil; i++ {
			next := cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}

		end, right := prev, cur
		start, end = end, start
		left.Next = start
		end.Next = right
	}

	return head
}
