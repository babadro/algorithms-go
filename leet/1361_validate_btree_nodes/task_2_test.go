package _1361_validate_btree_nodes

import "testing"

//n = 4, leftChild = [1,-1,3,-1], rightChild = [2,-1,-1,-1] true
// n = 4, leftChild = [1,-1,3,-1], rightChild = [2,3,-1,-1] false
// n = 2, leftChild = [1,0], rightChild = [-1,-1] false
//n = 6, leftChild = [1,-1,-1,4,-1,-1], rightChild = [2,-1,-1,5,-1,-1] false
func TestValidation(t *testing.T) {
	cases := []struct {
		n          int
		leftChild  []int
		rightChild []int
		expected   bool
	}{
		{4, []int{1, -1, 3, -1}, []int{2, -1, -1, -1}, true},
		{4, []int{1, -1, 3, -1}, []int{2, 3, -1, -1}, false},
		{2, []int{1, 0}, []int{-1, -1}, false},
		{6, []int{1, -1, -1, 4, -1, -1}, []int{2, -1, -1, 5, -1, -1}, false},
	}

	for i, c := range cases {
		if fact := validateBinaryTreeNodes(c.n, c.leftChild, c.rightChild); fact != c.expected {
			t.Errorf("case#%d, left: %v, right: %v, want %t, got %t", i+1, c.leftChild, c.rightChild, c.expected, fact)
		}
	}
}
