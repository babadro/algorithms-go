package _152_maximum_product_subarray

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{nums: []int{2, 3, -2, 4}}, want: 6},
		{name: "2", args: args{nums: []int{1}}, want: 1},
		{name: "3", args: args{nums: []int{0}}, want: 0},
		{name: "4", args: args{nums: []int{1, 1, 0}}, want: 1},
		{name: "5", args: args{nums: []int{-1, -2, -3, -4}}, want: 24},
		{name: "6", args: args{nums: []int{-2, 0, 0, 0, -2, -4, 2}}, want: 16},
		{name: "7", args: args{nums: []int{-2, 0, -1}}, want: 0},
		{name: "8", args: args{nums: []int{-2}}, want: -2},
		{name: "9", args: args{nums: []int{-2, 0}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
