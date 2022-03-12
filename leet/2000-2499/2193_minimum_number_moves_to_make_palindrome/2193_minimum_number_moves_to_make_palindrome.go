package _2193_minimum_number_moves_to_make_palindrome

import "container/list"

// passed. tptl. hard. need to understand
// https://leetcode.com/problems/minimum-number-of-moves-to-make-palindrome/discuss/1822174/C%2B%2BPython-Short-Greedy-Solution
func minMovesToMakePalindrome(s string) int {
	res := 0

	l := list.New()
	for i := range s {
		l.PushBack(s[i])
	}

	for l.Len() > 1 {
		it := l.Front()
		step := 0
		for ; it.Value.(byte) != l.Back().Value.(byte); it = it.Next() {
			step++
		}

		if step == l.Len()-1 {
			res += step / 2
		} else {
			res += step
			l.Remove(it)
		}

		l.Remove(l.Back())
	}

	return res
}

// todo 2 this is more effective solution:
func minMovesToMakePalindrome2(str string) int {
	s := []byte(str)

	left := 0
	right := len(s) - 1

	res := 0
	for left < right {
		last := right
		for s[last] != s[left] {
			last--
		}

		if last == left {
			res++
			s[last], s[last+1] = s[last+1], s[last]
		} else {
			for i := last; i < right; i++ {
				s[i], s[i+1] = s[i+1], s[i]
			}

			res += right - last
			right--
			left++
		}
	}

	return res
}
