package _543_tree_diameter

import (
	"testing"

	"github.com/babadro/algorithms-go/base/binaryTree"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	type args struct {
		root *binaryTree.Node
	}

	n1, n2, n3, n4, n5 := n(1), n(2), n(3), n(4), n(5)
	n1.Left, n1.Right = n2, n3
	n2.Left, n2.Right = n4, n5
	args1 := args{root: n1}

	n1, n2 = n(1), n(2)
	n1.Left = n2
	args2 := args{root: n1}

	tests := []struct {
		args args
		want int
	}{
		{args: args1, want: 3},
		{args2, 1},
		{
			args: args{
				root: constructTree([]interface{}{4, -7, -3, nil, nil, -9, -3, 9, -7, -4, nil, 6, nil, -6, -6, nil, nil, 0, 6, 5, nil, 9, nil, nil, -1, -4, nil, nil, nil, -2}),
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func n(val int) *binaryTree.Node {
	return &binaryTree.Node{Val: val}
}

func constructTree(nums []interface{}) *binaryTree.Node {
	if nums[0] == nil {
		return nil
	}

	root := &binaryTree.Node{Val: nums[0].(int)}
	queue := []*binaryTree.Node{root}
	i := 1
	for i < len(nums) {
		node := queue[0]
		queue = queue[1:]
		if nums[i] != nil {
			node.Left = &binaryTree.Node{Val: nums[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(nums) && nums[i] != nil {
			node.Right = &binaryTree.Node{Val: nums[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
