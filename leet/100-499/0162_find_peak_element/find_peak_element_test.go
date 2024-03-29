package _162_find_peak_element

import (
	"fmt"
	"testing"
)

func Test_findPeakElement(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 2},
		{[]int{1, 2, 1, 3, 5, 6, 4}, 5},
		{[]int{1}, 0},
		{[]int{1, 2, 3, 4}, 3},
		{[]int{4, 3, 2, 1}, 0},
		{[]int{1, 2}, 1},
		{[]int{2, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := findPeakElementIterative(tt.nums); got != tt.want {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindPeakElementIterative(t *testing.T) {
	cases := [][]int{
		{2, 1, 3, 5, 4},
		{1, 4, 3, 6, 5},
		{1, 2, 3, 4, 5, 6},
		{3, 2, 1},
	}

	for i, c := range cases {
		res := findPeakElementIterative(c)
		err := false
		if len(c) == 1 {
			if res != 0 {
				err = true
			}
		} else if res == 0 {
			if c[1] > c[res] {
				err = true
			}
		} else if res == len(c)-1 {
			if c[res-1] > c[res] {
				err = true
			}
		} else if c[res] < c[res+1] || c[res] < c[res-1] {
			err = true
		}
		if err {
			t.Errorf("case#%d, wrong answer %d", i+1, res)
		}
	}
}

func TestFindPeakElementRecursive(t *testing.T) {
	cases := [][]int{
		{2, 1, 3, 5, 4},
		{1, 4, 3, 6, 5},
		{1, 2, 3, 4, 5, 6},
		{3, 2, 1},
	}

	for i, c := range cases {
		res := findPeakElementRecursive(c)
		err := false
		if len(c) == 1 {
			if res != 0 {
				err = true
			}
		} else if res == 0 {
			if c[1] > c[res] {
				err = true
			}
		} else if res == len(c)-1 {
			if c[res-1] > c[res] {
				err = true
			}
		} else if c[res] < c[res+1] || c[res] < c[res-1] {
			err = true
		}
		if err {
			t.Errorf("case#%d, wrong answer %d", i+1, res)
		}
	}
}

func TestFindPeakElement(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 1, 3, 5, 4}, 0},
		{[]int{1, 4, 3, 6, 5}, 1},
		{[]int{1, 2, 3, 4, 5, 6}, 5},
		{[]int{3, 2, 1}, 0},
	}

	for i, c := range cases {
		if fact := findPeakElementRecursive(c.nums); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}
