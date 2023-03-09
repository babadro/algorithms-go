package _2074_reverse_nodes_in_even_length_groups

import (
	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/singly"
)

// tptl. passed
func reverseEvenLengthGroups(head *singly.ListNode) *singly.ListNode {
	if head == nil {
		return nil
	}

	cur, prev := head, head
	for groupSize := 1; cur != nil; groupSize++ {
		actualSize := 0
		tmpPrev, tmpCur := prev, cur
		for ; actualSize < groupSize && cur != nil; actualSize++ {
			prev = cur
			cur = cur.Next
		}

		if actualSize%2 == 1 {
			continue
		}

		prev, cur = tmpPrev, tmpCur

		left, start := prev, cur
		for i := 0; i < actualSize; i++ {
			next := cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}

		end, right := prev, cur
		start, end = end, start
		left.Next = start
		end.Next = right
		prev = end
	}

	return head
}
