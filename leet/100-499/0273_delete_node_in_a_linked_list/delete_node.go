package _273_delete_node_in_a_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	nodeToDelete := node.Next
	node.Val = nodeToDelete.Val
	node.Next = nodeToDelete.Next
	nodeToDelete.Next = nil
}
