package _450_delete_node_in_a_bst

import (
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"
	"log"
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	tests := []struct {
		name              string
		key               int
		inputTreeBuilder  func() *btree.Node
		resultTreeBuilder func() *btree.Node
	}{
		{
			name: "1",
			key:  3,
			inputTreeBuilder: func() *btree.Node {
				n5, n3, n6, n2, n4, n7 := n(5), n(3), n(6), n(2), n(4), n(7)

				n5.Left, n5.Right = n3, n6
				n3.Left, n3.Right = n2, n4
				n6.Left = n7

				return n5
			},
			resultTreeBuilder: func() *btree.Node {
				n5, n4, n6, n2, n7 := n(5), n(4), n(6), n(2), n(7)

				n5.Left, n5.Right = n4, n6
				n4.Left = n2
				n6.Left = n7

				return n5
			},
		},
		{
			name: "2",
			key:  0,
			inputTreeBuilder: func() *btree.Node {
				n5, n2, n6, n4, n7 := n(5), n(2), n(6), n(4), n(7)
				n5.Left, n5.Right = n2, n6
				n2.Right = n4
				n6.Right = n7

				return n5
			},
			resultTreeBuilder: func() *btree.Node {
				n5, n2, n6, n4, n7 := n(5), n(2), n(6), n(4), n(7)
				n5.Left, n5.Right = n2, n6
				n2.Right = n4
				n6.Right = n7

				return n5
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, want := tt.inputTreeBuilder(), tt.resultTreeBuilder()
			got := deleteNode(input, tt.key)
			if !reflect.DeepEqual(got, want) {
				t.Error("Wrong result")
				t.Log("got: ")
				btree.InOrderFunc(got, func(node *btree.Node) {
					log.Println(node.Val)
				})

				t.Log("want: ")
				btree.InOrderFunc(want, func(node *btree.Node) {
					log.Println(node.Val)
				})
			}
		})
	}
}

func n(val int) *btree.Node {
	return &btree.Node{Val: val}
}
