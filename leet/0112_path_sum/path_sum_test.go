package _112_path_sum

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		sum  int
		want bool
	}{
		{"1", []int{5, 4, 8, 11, binaryTree.Null, 13, 4, 7, 2, binaryTree.Null, binaryTree.Null, 1}, 22, true},
		{"2", []int{1, 2, 3}, 3, true},
		{"3", []int{1, 2, 3}, 4, true},
		{"4", []int{1, 2, 3}, 1, false},
		{"5", []int{1, 2, binaryTree.Null, 3}, 6, true},
		{"6", []int{2}, 2, true},
		{"7", []int{}, 2, false},
		{"9", []int{1, 2}, 3, true},
		{"10", []int{1, 2, binaryTree.Null, binaryTree.Null, 3}, 3, false},
		{"11", []int{-2, binaryTree.Null, -3}, -5, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := binaryTree.ArrayToBinaryTree(tt.arr)
			if got := hasPathSum2(root, tt.sum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
