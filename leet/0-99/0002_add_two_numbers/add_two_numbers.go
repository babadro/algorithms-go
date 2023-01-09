package _002_add_two_numbers

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

func addTwoNumbers(l1 *single.ListNode, l2 *single.ListNode) *single.ListNode {
	var head, prev *single.ListNode
	node1, node2 := l1, l2
	carry := false
	for node1 != nil || node2 != nil || carry {
		sum := 0
		if node1 != nil {
			sum += node1.Val
			node1 = node1.Next
		}

		if node2 != nil {
			sum += node2.Val
			node2 = node2.Next
		}

		if carry {
			sum += 1
		}

		if sum > 9 {
			sum %= 10
			carry = true
		} else {
			carry = false
		}

		if head == nil {
			prev = &single.ListNode{Val: sum}
			head = prev
		} else {
			prev.Next = &single.ListNode{Val: sum}
			prev = prev.Next
		}
	}

	return head
}
