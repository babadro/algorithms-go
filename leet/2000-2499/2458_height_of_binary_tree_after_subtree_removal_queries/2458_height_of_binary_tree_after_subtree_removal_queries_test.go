package _2458_height_of_binary_tree_after_subtree_removal_queries

import (
	"reflect"
	"testing"

	"github.com/babadro/algorithms-go/base/binaryTree"
	"github.com/babadro/algorithms-go/utils"
	"github.com/stretchr/testify/require"
)

func Test_treeQueries(t *testing.T) {
	tests := []struct {
		root    *binaryTree.Node
		queries []int
		want    []int
	}{
		{binaryTree.InterfacesArrToBinaryTree(
			1, 3, 4, 2, nil, 6, 5, nil, nil, nil, nil, nil, 7), []int{4}, []int{2},
		},
		{
			binaryTree.InterfacesArrToBinaryTree(5, 8, 9, 2, 1, 3, 7, 4, 6), []int{3, 2, 4, 8},
			[]int{3, 2, 3, 2},
		},
		{
			binaryTree.InterfacesArrToBinaryTree(
				utils.ParseNullableNums[int]("testdata/tle_tree.txt")...),
			utils.ParseNums[int]("testdata/tle_queries.txt"), utils.ParseNums[int]("testdata/tle_answer.txt"),
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := treeQueries2(tt.root, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("treeQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_height(t *testing.T) {
	tests := []struct {
		node *binaryTree.Node
		want []pathNode
	}{
		{
			binaryTree.InterfacesArrToBinaryTree(1, 3, nil, 2), []pathNode{
				{val: 2, height: 1}, {val: 3, height: 2}, {val: 1, height: 3},
			},
		},
		{
			binaryTree.InterfacesArrToBinaryTree(1, 3, 4, 2, nil, 6, 5, nil, nil, nil, nil, nil, 7), []pathNode{
				{7, 1, 0}, {5, 2, 0}, {4, 3, 1}, {1, 4, 2},
			},
		},
		{
			binaryTree.InterfacesArrToBinaryTree(1, 3, 4, 2, nil, 6, 5, nil, nil, nil, 9, nil, 7), []pathNode{
				{4, 3, 2}, {1, 4, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := heights(tt.node)
			require.Equal(t, tt.want, got)
		})
	}
}
