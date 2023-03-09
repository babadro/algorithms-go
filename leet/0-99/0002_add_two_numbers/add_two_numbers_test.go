package _002_add_two_numbers

import (
	"strconv"
	"testing"

	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/singly"
)

func TestAddTwoNumbers(t *testing.T) {
	cases := []struct {
		term1    string
		term2    string
		expected string
	}{
		{term1: "0", term2: "0", expected: "0"},
		{"0", "2", "2"},
		{"3", "0", "3"},
		{"99", "0", "99"},
		{"0", "83", "83"},
		{"903", "217", "1120"},
		{"99", "85", "184"},
		{"100000000000000000000000000000000000000000000000000000000000001", "10", "100000000000000000000000000000000000000000000000000000000000011"},
	}
	for i, c := range cases {
		l1, l2 := StringToLinkedList(c.term1), StringToLinkedList(c.term2)
		if fact := LinkedListToString(addTwoNumbers(l1, l2)); fact != c.expected {
			t.Errorf("case:#%d, expected %s, got %s", i+1, c.expected, fact)
		}
	}
}

func TestLinkedListFromString(t *testing.T) {
	cases := []string{"0", "1", "2", "100", "992"}
	for _, c := range cases {
		node := StringToLinkedList(c)
		t.Log(LinkedListToString(node))
	}
}

func StringToLinkedList(input string) *singly.ListNode {
	res := &singly.ListNode{Val: int(input[len(input)-1]) - 48}
	prevNode := res
	for i := len(input) - 2; i >= 0; i-- {
		newNode := &singly.ListNode{Val: int(input[i] - 48)}
		prevNode.Next = newNode
		prevNode = newNode
	}
	return res
}

func LinkedListToString(node *singly.ListNode) string {
	res := make([]byte, 0)
	for node != nil {
		res = append(res, strconv.Itoa(node.Val)...)
		node = node.Next
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return string(res)
}
