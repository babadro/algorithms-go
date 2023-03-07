package _092_reverse_linked_list_ii

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

// tptl. passed. medium #linkedlist
func reverseBetween(head *single.ListNode, left int, right int) *single.ListNode {
	leftNode := head
	var prevLeft *single.ListNode
	for i := 1; i < left; i++ {
		prevLeft, leftNode = leftNode, leftNode.Next
	}

	// reverse
	var prev *single.ListNode
	cur := leftNode
	for i := left; i <= right; i++ {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}

	rightNode, nextRight := prev, cur
	leftNode, rightNode = rightNode, leftNode

	if prevLeft != nil {
		prevLeft.Next = leftNode
	} else {
		head = leftNode
	}

	rightNode.Next = nextRight

	return head
}

// looks more complicated
func reverseBetween2(head *single.ListNode, left int, right int) *single.ListNode {
	var leftNode, rightNode, prevLeft, nextRight, prev *single.ListNode
	for cur, i := head, 1; ; i++ {
		if i == left {
			leftNode, prevLeft = cur, prev
		}

		next := cur.Next

		if i > left { // reverse
			cur.Next = prev
		}

		if i == right {
			rightNode, nextRight = cur, next

			break
		}

		prev, cur = cur, next
	}

	if prevLeft != nil {
		prevLeft.Next = rightNode
	} else {
		head = rightNode
	}

	leftNode.Next = nextRight

	return head
}

// slower, but easy to understand
func reverseBetween3(head *single.ListNode, left int, right int) *single.ListNode {
	var headSegmentEnd, tailSegmentStart, leftNode, rightNode *single.ListNode

	for i, cur, prev := 1, head, (*single.ListNode)(nil); cur != nil; {
		if i == left {
			headSegmentEnd = prev
			leftNode = cur
		}

		if i == right {
			tailSegmentStart = cur.Next
			rightNode = cur

			break
		}

		i++
		prev, cur = cur, cur.Next
	}

	rightNode.Next = nil

	reversed := reverse(leftNode)
	if headSegmentEnd != nil {
		headSegmentEnd.Next = reversed
	} else {
		head = reversed
	}

	leftNode.Next = tailSegmentStart

	return head
}

func reverse(node *single.ListNode) *single.ListNode {
	var prev *single.ListNode
	for node != nil {
		next := node.Next
		node.Next = prev
		prev, node = node, next
	}

	return prev
}
