package _234_palindrome_linked_list

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

// tptl. passed. easy to understand
func isPalindrome(head *single.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	tail := reverse(slow)
	node1, node2 := head, tail

	palindrome := true
	for node1 != nil && node2 != nil {
		if node1.Val != node2.Val {
			palindrome = false
			break
		}

		node1, node2 = node1.Next, node2.Next
	}

	reverse(tail)

	return palindrome
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

// with slice allocation
func isPalindrome3(head *single.ListNode) bool {
	var arr []int
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next

		arr = append(arr, slow.Val)
		slow = slow.Next
	}

	if len(arr) == 0 {
		return true
	}

	if fast != nil { // list len is odd
		slow = slow.Next
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != slow.Val {
			return false
		}

		slow = slow.Next
	}

	return true
}
