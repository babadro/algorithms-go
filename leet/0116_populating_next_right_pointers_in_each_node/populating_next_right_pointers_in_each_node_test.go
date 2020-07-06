package _116_populating_next_right_pointers_in_each_node

import (
	"testing"
)

func TestConnect(t *testing.T) {
	//n8, n9, n10, n11, n12, n13, n14, n15 := &Node{Val: 8}, &Node{Val: 9}, &Node{Val: 10}, &Node{Val: 11}, &Node{Val: 12}, &Node{Val: 13}, &Node{Val: 14}, &Node{Val: 15}
	//n4, n5, n6, n7 := &Node{Val: 4}, &Node{Val: 5}, &Node{Val: 6}, &Node{Val: 7}
	n2, n3 := &Node{Val: 2}, &Node{Val: 3}
	n1 := &Node{Val: 1}

	n1.Left, n1.Right = n2, n3
	//n2.Left, n2.Right, n3.Left, n3.Right = n4, n5, n6, n7
	//n4.Left, n4.Right, n5.Left, n5.Right, n6.Left, n6.Right, n7.Left, n7.Right = n8, n9, n10, n11, n12, n13, n14, n15
	_ = n1
	connect(n1)
	//root := &Node{Val: 1}
	//connect(root)
}
