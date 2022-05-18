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

func isPalindromeTheBest(head *single.ListNode) bool {
	slow, fast := head, head
	var prev *single.ListNode
	for fast != nil && fast.Next != nil {
		cur := slow
		slow, fast = slow.Next, fast.Next.Next
		// reverse first half of linked list
		cur.Next, prev = prev, cur
		// alternative implementation - todo 3 how does it work?
		//prev, slow.Next, slow, fast = slow, prev, slow.Next, fast.Next.Next
	}

	if fast != nil { // len(list) is odd
		slow = slow.Next
	}

	palindrome := true
	for prev != nil {
		if prev.Val != slow.Val {
			palindrome = false
		}

		// reverse again to reassemble initial order
		prev, slow = prev.Next, slow.Next
	}

	return palindrome
}
