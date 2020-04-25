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

func LinkedListToArr(head *ListNode) (arr []int) {
	node := head
	for node != nil {
		arr = append(arr, node.Val)
		node = node.Next
	}
	return arr
}
