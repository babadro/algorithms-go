package _279_perfect_squares

import (
	"algorithms-go/slices"
	"testing"
)

func TestPerfectSquareNums(t *testing.T) {
	cases := []struct {
		upperBound         int
		expectedSquareNums []int
	}{
		{1, []int{1}},
		{2, []int{1}},
		{3, []int{1}},
		{4, []int{1, 4}},
		{5, []int{1, 4}},
		{9, []int{1, 4, 9}},
		{16, []int{1, 4, 9, 16}},
		{19, []int{1, 4, 9, 16}},
		{20, []int{1, 4, 9, 16}},
		{21, []int{1, 4, 9, 16}},
		{22, []int{1, 4, 9, 16}},
		{23, []int{1, 4, 9, 16}},
		{24, []int{1, 4, 9, 16}},
		{25, []int{1, 4, 9, 16, 25}},
	}

	for i, c := range cases {
		if fact := perfectSquareNums(c.upperBound); !slices.IntSlicesAreEqual(fact, c.expectedSquareNums) {
			t.Errorf("case#%d got %v, want %v", i+1, fact, c.expectedSquareNums)
		}
	}

}
