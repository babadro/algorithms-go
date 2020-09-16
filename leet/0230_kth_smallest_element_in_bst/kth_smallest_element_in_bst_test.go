package _230_kth_smallest_element_in_bst

import (
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{[]int{3, 1, 4, btree.Null, 2}, 1}, want: 1},
		{name: "2", args: args{[]int{5, 3, 6, 2, 4, btree.Null, btree.Null, 1}, 3}, want: 3},
		{name: "3", args: args{[]int{}, 1}, want: 0},
		{name: "4", args: args{[]int{1}, 1}, want: 1},
		{name: "4", args: args{[]int{2, 1}, 2}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := btree.ArrayToBinaryTree(tt.args.arr)
			if got := kthSmallest(root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
