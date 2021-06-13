package _45

import "testing"

func Test_maximumRemovals(t *testing.T) {
	tests := []struct {
		s         string
		p         string
		removable []int
		want      int
	}{
		{"abcacb", "ab", []int{3, 1, 0}, 2},
		{"abcbddddd", "abcd", []int{3, 2, 1, 4, 5, 6}, 1},
		{"abcab", "abc", []int{0, 1, 2, 3, 4}, 0},
		{"qlevcvgzfpryiqlwy", "qlecfqlw", []int{12, 5}, 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maximumRemovals(tt.s, tt.p, tt.removable); got != tt.want {
				t.Errorf("maximumRemovals() = %v, want %v", got, tt.want)
			}
		})
	}
}
