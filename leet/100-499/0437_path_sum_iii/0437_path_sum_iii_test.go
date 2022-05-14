package _437_path_sum_iii

import (
	"testing"

	"github.com/babadro/algorithms-go/base/binaryTree"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *binaryTree.Node
		targetSum int
	}

	n1_10 := n(10)
	n2_5, n2_m3 := n(5), n(-3)
	n3_3, n3_2, n3_11 := n(3), n(2), n(11)
	n4_3, n4_m2, n4_1 := n(3), n(-2), n(1)

	n1_10.Left, n1_10.Right = n2_5, n2_m3
	n2_5.Left, n2_5.Right = n3_3, n3_2
	n2_m3.Right = n3_11
	n3_3.Left, n3_3.Right = n4_3, n4_m2
	n3_2.Right = n4_1

	args1 := args{
		root:      n1_10,
		targetSum: 8,
	}

	n1_5 := n(5)
	n2_4, n2_8 := n(4), n(8)
	n3_11, n3_13, n3_4 := n(11), n(13), n(4)
	n4_7, n4_2, n4_5, n4_1 := n(7), n(2), n(5), n(1)

	n1_5.Left, n1_5.Right = n2_4, n2_8
	n2_4.Left = n3_11
	n2_8.Left, n2_8.Right = n3_13, n3_4
	n3_11.Left, n3_11.Right = n4_7, n4_2
	n3_4.Left, n3_4.Right = n4_5, n4_1

	args2 := args{
		root:      n1_5,
		targetSum: 22,
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args1, want: 3},
		{args: args2, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func n(v int) *binaryTree.Node {
	return &binaryTree.Node{Val: v}
}
