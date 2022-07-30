package _143_reorder_list

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

// tptl. best solution. fast
func reorderList(head *single.ListNode) {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	secondHead := slow.Next
	slow.Next = nil

	// reverse second half
	prev, cur := (*single.ListNode)(nil), secondHead
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}

	// reorder
	for cur1, cur2 := head, prev; cur1 != nil && cur2 != nil; {
		next1, next2 := cur1.Next, cur2.Next
		cur1.Next = cur2
		cur2.Next = next1
		cur1, cur2 = next1, next2
	}
}

// bruteforce. passed, but slow
func reorderList2(head *single.ListNode) {
	cur, prevTail, tail := head, (*single.ListNode)(nil), head
	for {
		for tail.Next != nil {
			prevTail, tail = tail, tail.Next
		}

		if tail == cur || prevTail == cur {
			break
		}

		next := cur.Next
		cur.Next = tail
		prevTail.Next = nil
		cur = next
		tail.Next = cur
	}
}
