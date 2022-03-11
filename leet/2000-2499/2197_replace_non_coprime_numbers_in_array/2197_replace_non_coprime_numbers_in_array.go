package _2197_replace_non_coprime_numbers_in_array

import (
	"container/list"
)

// tptl. passed. hard
// https://leetcode.com/problems/replace-non-coprime-numbers-in-array/discuss/1823681/List
// todo 2: slow solution (28%, 6%). Find quicker solution
func replaceNonCoprimes(nums []int) []int {
	l := list.New()
	for _, num := range nums {
		l.PushBack(num)
	}

	for it := l.Front(); it != nil; it = it.Next() {
		if it != l.Front() {
			n1, n2 := it.Prev().Value.(int), it.Value.(int)
			itGcd := gcd(n1, n2)
			if itGcd > 1 {
				l.Remove(it.Prev())
				it.Value = n1 * n2 / itGcd // LCM (Least Common Multiple)
				if it != l.Front() {
					it = it.Prev()
				}
			}
		}
	}

	res := make([]int, 0, l.Len())
	for it := l.Front(); it != nil; it = it.Next() {
		res = append(res, it.Value.(int))
	}

	return res
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}
