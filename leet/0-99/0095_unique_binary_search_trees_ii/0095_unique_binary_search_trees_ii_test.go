package _095_unique_binary_search_trees_ii

import (
	"sort"
	"strconv"
	"testing"

	"github.com/babadro/algorithms-go/base/binaryTree"
	"github.com/stretchr/testify/require"
)

func Test_generateTrees(t *testing.T) {
	res := make([]*binaryTree.Node, 5)
	n1, n2, n3 := n(1), n(2), n(3)
	n1.Right = n3
	n3.Left = n2
	res[0] = n1

	n1, n2, n3 = n(1), n(2), n(3)
	n1.Right, n2.Right = n2, n3
	res[1] = n1

	n1, n2, n3 = n(1), n(2), n(3)
	n2.Left, n2.Right = n1, n3
	res[2] = n1

	n1, n2, n3 = n(1), n(2), n(3)
	n3.Left, n2.Left = n2, n1
	res[3] = n1

	n1, n2, n3 = n(1), n(2), n(3)
	n3.Left, n1.Right = n1, n2
	res[4] = n1

	tests := []struct {
		n    int
		want []*binaryTree.Node
	}{
		{3, res},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			got := generateTrees(tt.n)
			sortRoots := func(roots []*binaryTree.Node) {
				sort.Slice(roots, func(i, j int) bool {
					return roots[i].Val < roots[j].Val
				})
			}

			sortRoots(got)
			sortRoots(tt.want)

			require.Equal(t, len(tt.want), len(got))
			for i := range got {
				var arrGot, arrWant []int
				binaryTree.InOrder(got[i], &arrGot)
				binaryTree.InOrder(tt.want[i], &arrWant)

				require.Equal(t, arrWant, arrGot)
			}
		})
	}
}

func n(v int) *binaryTree.Node {
	return &binaryTree.Node{Val: v}
}
