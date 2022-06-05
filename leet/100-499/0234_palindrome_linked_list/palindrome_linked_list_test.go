package _234_palindrome_linked_list

import (
	"testing"

	"github.com/babadro/algorithms-go/03_StacksAndQueues/04_LinkedList/single"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		input    []int
		expected bool
	}{
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{1, 1}, true},
		{[]int{1, 2}, false},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 2, 3, 4, 3, 2, 1}, true},
		{[]int{1, 2, 3, 3, 2, 1}, true},
	}

	for i, c := range cases {
		head := single.ArrToLinkedList(c.input)
		if fact := isPalindrome(head); fact != c.expected {
			t.Errorf("case#%d, want %t, got %t", i+1, c.expected, fact)
		}
		a := 1
		_ = a
	}
}
