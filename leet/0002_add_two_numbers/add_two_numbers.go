package _002_add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	mem := 0
	res := &ListNode{}
	var prevNode *ListNode
	val1, val2 := 0, 0
	val1 = l1.Val
	val2 = l2.Val
	sum := val1 + val2
	if sum < 10 {
		res.Val = sum
	} else {
		res.Val = sum % 10
		mem = 1
	}
	prevNode = res
	l1, l2 = l1.Next, l2.Next
	for l1 != nil || l2 != nil || mem != 0 {
		newNode := &ListNode{}
		prevNode.Next = newNode
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		sum := val1 + val2 + mem
		if sum < 10 {
			newNode.Val = sum
			mem = 0
		} else {
			newNode.Val = sum % 10
			mem = 1
		}
		prevNode = newNode
	}
	return res
}
