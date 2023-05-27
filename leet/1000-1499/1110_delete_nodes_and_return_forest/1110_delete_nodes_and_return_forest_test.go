package _1110_delete_nodes_and_return_forest

import (
	"sort"
	"testing"

	"github.com/babadro/algorithms-go/base/binaryTree"
	"github.com/stretchr/testify/require"
)

func Test_delNodes(t *testing.T) {
	tests := []struct {
		root      *binaryTree.Node
		to_delete []int
		want      []*binaryTree.Node
	}{
		{
			binaryTree.InterfacesArrToBinaryTree(1, 2, 3, 4, 5, 6, 7),
			[]int{3, 5},
			[]*binaryTree.Node{
				binaryTree.InterfacesArrToBinaryTree(1, 2, nil, 4),
				{Val: 6},
				{Val: 7},
			},
		},
		{
			binaryTree.InterfacesArrToBinaryTree(1, 2, 4, nil, 3),
			[]int{3},
			[]*binaryTree.Node{
				binaryTree.InterfacesArrToBinaryTree(1, 2, 4),
			},
		},
		{
			binaryTree.InterfacesArrToBinaryTree(1, 2, 3, nil, nil, nil, 4),
			[]int{2, 1},
			[]*binaryTree.Node{
				binaryTree.InterfacesArrToBinaryTree(3, nil, 4),
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := delNodes(tt.root, tt.to_delete)

			require.Equal(t, len(tt.want), len(got))
			sortNodes(tt.want)
			sortNodes(got)

			require.Equal(t, tt.want, got)
		})
	}
}

func sortNodes(nodes []*binaryTree.Node) {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})
}
