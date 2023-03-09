package _025_reverse_nodes_in_k_group

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

// tptl. passed. hard
func reverseKGroup(head *single.ListNode, k int) *single.ListNode {
	var prev *single.ListNode
	for cur := head; cur != nil; {
		reversed, next, ok := reverse(cur, k)
		if !ok {
			reversed, next, _ = reverse(reversed, k)
		}

		if prev != nil {
			prev.Next = reversed
		} else {
			head = reversed
		}

		prev, cur = cur, next
	}

	return head
}

func reverse(node *single.ListNode, k int) (head, next *single.ListNode, ok bool) {
	var prev *single.ListNode
	for ; node != nil && k > 0; k-- {
		next = node.Next
		node.Next = prev
		prev, node = node, next
	}

	return prev, next, k == 0
}

// longer and slower
func reverseKGroup2(head *single.ListNode, k int) *single.ListNode {
	if k == 1 {
		return head
	}

	var l, r *single.ListNode
	start, end := head, head
	for {
		for count := 1; count < k; count++ {
			if end == nil || end.Next == nil {
				return head
			}

			end = end.Next
		}

		r = end.Next

		prev, cur := l, start
		for cur != r {
			next := cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}

		start, end = end, start

		end.Next = r
		if l != nil {
			l.Next = start
		} else {
			head = start
		}

		l = end
		start, end = end.Next, end.Next
	}
}
