package single

type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrToLinkedList(arr []int) *ListNode {
	var node *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		newNode := &ListNode{Val: arr[i], Next: node}
		node = newNode
	}
	return node
}
