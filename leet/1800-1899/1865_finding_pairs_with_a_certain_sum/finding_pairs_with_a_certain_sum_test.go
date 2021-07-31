package _1865_finding_pairs_with_a_certain_sum

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	s := Constructor([]int{1, 1, 2, 2, 2, 3}, []int{1, 4, 5, 2, 5, 4})
	t.Log(s.Count(7))
	s.Add(3, 2)
	t.Log(s.Count(8))
	t.Log(s.Count(4))
	s.Add(0, 1)
	s.Add(1, 1)
	t.Log(s.Count(7))
}
