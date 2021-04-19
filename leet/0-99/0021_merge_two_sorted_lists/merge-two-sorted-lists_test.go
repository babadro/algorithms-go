package _021_merge_two_sorted_lists

import (
	"fmt"
	"testing"
)

//1->2->4, 1->3->4
func TestMergeTwoSortedLists(t *testing.T) {
	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 3}
	head1.Next.Next = &ListNode{Val: 5}

	head2 := &ListNode{Val: 78}
	head2.Next = &ListNode{Val: 78}
	head2.Next.Next = &ListNode{Val: 78}

	mergedList := mergeTwoLists(head1, head2)
	printList(mergedList)
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
