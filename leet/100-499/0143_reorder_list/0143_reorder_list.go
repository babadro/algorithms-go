package _143_reorder_list

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

// tptl. best solution. fast
func reorderList(head *single.ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	left, right := head, reverse(slow)
	for left != nil && right != nil && left.Next != right {
		leftNext, rightNext := left.Next, right.Next
		left.Next = right
		right.Next = leftNext

		left, right = leftNext, rightNext
	}
}

func reverse(node *single.ListNode) *single.ListNode {
	var prev *single.ListNode
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}

	return prev
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
