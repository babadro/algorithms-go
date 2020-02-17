package _1339_maximum_product_of_splitted_binary_tree

/*
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func calcProduct(root1, root2 *TreeNode) int {
	var sum1 int
	f := func(node *TreeNode) { sum1 += node.Val }
	InOrderFunc(root1, f)

	var sum2 int
	f = func(node *TreeNode) { sum2 += node.Val }
	InOrderFunc(root2, f)

	return sum1 * sum2
}

func maxProduct(root *TreeNode) int {
	var product int
	f := func(x *TreeNode) {
		if x.Left != nil {
			root1, root2 := root, x.Left
			x.Left = nil
			if prod := calcProduct(root1, root2); prod > product {
				product = prod
			}
			x.Left = root2
		}
		if x.Right != nil {
			root1, root2 := root, x.Right
			x.Right = nil
			if prod := calcProduct(root1, root2); prod > product {
				product = prod
			}
			x.Right = root2
		}
	}
	InOrderFunc(root, f)
	return product % (int(1e9) + 7)
}

func InOrderFunc(node *TreeNode, f func(node *TreeNode)) {
	if node != nil {
		InOrderFunc(node.Left, f)
		f(node)
		InOrderFunc(node.Right, f)
	}
}
*/
