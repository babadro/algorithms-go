package _279_perfect_squares

import (
	"github.com/babadro/algorithms-go/slices"
	"testing"
)

func TestNumSquares(t *testing.T) {
	cases := []struct {
		num      int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 1},
		{5, 2},
		{12, 3},
		{13, 2},
	}

	for i, c := range cases {
		if fact := numSquares(c.num); fact != c.expected {
			t.Errorf("case#%d: want %d, got %d", i+1, c.expected, fact)
		}
	}
}

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

func TestCoinChanging(t *testing.T) {
	arr := []int{1, 4, 9}
	amount := 12
	t.Log(DynamicCoinChanging(arr, amount))
}
