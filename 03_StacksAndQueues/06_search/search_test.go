package _6_search

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{nil, 1}, -1},
		{"1 ok", args{[]int{1}, 1}, 0},
		{"1 fail", args{[]int{1}, 2}, -1},
		{"2 ok 1", args{[]int{1, 3}, 3}, 1},
		{"2 ok 2", args{[]int{1, 3}, 1}, 0},
		{"2 fail 1", args{[]int{1, 3}, 2}, -1},
		{"2 fail 2", args{[]int{1, 3}, 0}, -1},
		{"2 fail 3", args{[]int{1, 3}, 4}, -1},
		{"2 fail 4", args{[]int{1, 3}, 2}, -1},
		{"", args{[]int{1, 3, 4, 5}, 3}, 1},
		{"", args{[]int{-1, 0, 3, 5, 9, 12}, 9}, 4},
		{"", args{[]int{-1, 0, 3, 5, 9, 12}, 2}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
