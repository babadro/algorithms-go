package _637_average_of_levels_in_binary_tree

import (
	"reflect"
	"testing"

	"github.com/babadro/algorithms-go/base/binaryTree"
)

func Test_averageOfLevels(t *testing.T) {
	type args struct {
		root *binaryTree.Node
	}

	n3, n9, n20, n15, n7 := n(3), n(9), n(20), n(15), n(7)
	n3.Left, n3.Right = n9, n20
	n20.Left, n20.Right = n15, n7

	args1 := args{root: n3}

	n3, n9, n20, n15, n7 = n(3), n(9), n(20), n(15), n(7)
	n3.Left, n3.Right = n9, n20
	n9.Left, n9.Right = n15, n7

	args2 := args{root: n3}

	tests := []struct {
		args args
		want []float64
	}{
		{args1, []float64{3, 14.5, 11}},
		{args2, []float64{3, 14.5, 11}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := averageOfLevels(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("averageOfLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func n(v int) *binaryTree.Node {
	return &binaryTree.Node{Val: v}
}
