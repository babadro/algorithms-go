package _1629_slowest_key

import "testing"

func Test_slowestKey(t *testing.T) {
	tests := []struct {
		releaseTimes []int
		keysPressed  string
		want         byte
	}{
		{[]int{9, 29, 49, 50}, "cbcd", 'c'},
		{[]int{12, 23, 36, 46, 62}, "spuda", 'a'},
		{[]int{1, 2, 3}, "bbb", 'b'},
		{[]int{1, 2, 3}, "cba", 'c'},
		{[]int{1, 3}, "aa", 'a'},
	}
	for _, tt := range tests {
		t.Run(tt.keysPressed, func(t *testing.T) {
			if got := slowestKey2(tt.releaseTimes, tt.keysPressed); got != tt.want {
				t.Errorf("slowestKey() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}
