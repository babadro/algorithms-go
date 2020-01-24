package _002_add_two_numbers

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var str1, str2 []byte
	node := l1
	for node != nil {
		str1 = append(str1, strconv.Itoa(node.Val)...)
		node = node.Next
	}

	//for i := 0; i < len(str1)/2; i++ {
	//	str1[i], str1[len(str1)-1-i] = str1[len(str1)-1-i], str1[i]
	//}
	node = l2
	for node != nil {
		str2 = append(str2, strconv.Itoa(node.Val)...)
		node = node.Next
	}

	var int1, int2 int

	int1, _ = strconv.Atoi(string(str1))
	int2, _ = strconv.Atoi(string(str2))
	sum := int1 + int2
	sumStr := []byte(strconv.Itoa(sum))
	//for i := 0; i < len(sumStr)/2; i++ {
	//	sumStr[i], sumStr[len(str1)-1-i] = sumStr[len(str1)-1-i], sumStr[i]
	//}
	firstDigit, _ := strconv.Atoi(string(sumStr[0]))
	node = &ListNode{Val: firstDigit}
	for i := 1; i < len(sumStr); i++ {
		digit, _ := strconv.Atoi(string(sumStr[i]))
		newNode := &ListNode{Val: digit, Next: node}
		node = newNode
	}
	return node
}
