package _2074_reverse_nodes_in_even_length_groups

import (
	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/singly"
)

// best solution, one pass
func reverseEvenLengthGroups2(head *singly.ListNode) *singly.ListNode {
	cur := head

	expectedLen := 1
	var prev *singly.ListNode
	for ; cur != nil; expectedLen++ {
		if expectedLen%2 == 1 {
			tmpPrev, tmpCur, actualLen := prev, cur, 0
			for ; tmpCur != nil && actualLen < expectedLen; actualLen++ {
				tmpPrev, tmpCur = tmpCur, tmpCur.Next
			}

			if actualLen%2 == 1 {
				prev, cur = tmpPrev, tmpCur
				continue
			}
		}

		reversed, next, ok := reverse(cur, expectedLen)
		if !ok {
			_, _, _ = reverse(reversed, expectedLen)
			return head
		}

		cur.Next = next
		prev.Next = reversed
		prev, cur = cur, next
	}

	return head
}

func reverse(node *singly.ListNode, k int) (*singly.ListNode, *singly.ListNode, bool) {
	var prev *singly.ListNode
	actualLen := 0
	for ; node != nil && k > 0; k-- {
		next := node.Next
		node.Next = prev
		prev, node = node, next

		actualLen++
	}

	return prev, node, actualLen%2 == 0
}

// tptl. passed. 2 passes
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
