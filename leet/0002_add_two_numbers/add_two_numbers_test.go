package _002_add_two_numbers

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	//Output: 7 -> 0 -> 8
	l1 := &ListNode{Val: 3}
	l2 := &ListNode{Val: 0}
	l3 := &ListNode{Val: 9}

	l1.Next = l2
	l2.Next = l3

	node := l1
	for node != nil {
		t.Log(node.Val)
		node = node.Next
	}

	l4 := &ListNode{Val: 7}
	l5 := &ListNode{Val: 1}
	l6 := &ListNode{Val: 2}
	l4.Next = l5
	l5.Next = l6

	node = l4
	for node != nil {
		t.Log(node.Val)
		node = node.Next
	}
	node = addTwoNumbers(l4, l1)
	for node != nil {
		t.Log(node.Val)
		node = node.Next
	}

	l1 = &ListNode{Val: 0}

	l2 = &ListNode{Val: 7}
	l3 = &ListNode{Val: 3}

	l2.Next = l3
	node = addTwoNumbers(l1, l2)
	for node != nil {
		t.Log(node.Val)
		node = node.Next
	}

	l1 = &ListNode{Val: 1}
	l2 = &ListNode{Val: 8}
	l1.Next = l2

	l3 = &ListNode{Val: 0}
	node = addTwoNumbers(l1, l3)
	for node != nil {
		t.Log(node.Val)
		node = node.Next
	}
}
