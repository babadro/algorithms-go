package _234_palindrome_linked_list

import "github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"

func isPalindrome(head *single.ListNode) bool {
	totalCount := 0
	node := head
	for node != nil {
		totalCount++
		node = node.Next
	}
	length := totalCount / 2
	var nodeToReverse *single.ListNode
	counter := 0
	nodeToReverse = head
	for {
		if counter == totalCount-length {
			break
		}
		nodeToReverse = nodeToReverse.Next
		counter++

	}
	node1, node2 := head, reverse(nodeToReverse)
	for i := 0; i < length; i++ {
		if node1.Val != node2.Val {
			return false
		}
		node1 = node1.Next
		node2 = node2.Next
	}
	return true
}

func reverse(head *single.ListNode) *single.ListNode {
	node := head
	var prev *single.ListNode
	var next *single.ListNode
	for node != nil {
		next = node.Next
		node.Next = prev
		prev = node
		node = next
	}
	return prev
}

// TODO 3 need to understand
func isPalindromeTheBest(head *single.ListNode) bool {
	head, slow, fast := nil, head, head
	for fast != nil && fast.Next != nil {
		head, slow.Next, slow, fast = slow, head, slow.Next, fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	for ; head != nil && head.Val == slow.Val; head, slow = head.Next, slow.Next {
	}
	return head == nil
}
