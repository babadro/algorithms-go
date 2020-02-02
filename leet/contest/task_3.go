package contest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderSplit(x, root *TreeNode, product *int) {
	if x != nil {
		InorderSplit(x.Left, root, product)
		if x.Left != nil {
			root1, root2 := root, x.Left
			x.Left = nil
			if prod := calcProduct(root1, root2); prod > *product {
				*product = prod
			}
			x.Left = root2
		}
		if x.Right != nil {
			root1, root2 := root, x.Right
			x.Right = nil
			if prod := calcProduct(root1, root2); prod > *product {
				*product = prod
			}
			x.Right = root2
		}
		InorderSplit(x.Right, root, product)
	}
}

func calcProduct(root1, root2 *TreeNode) int {
	sum1 := 0
	InorderSum(root1, &sum1)

	sum2 := 0
	InorderSum(root2, &sum2)

	i := 1
	m := 1000000007
	i = (i * sum1) % m
	i = (i * sum2) % m
	return i
}

func InorderSum(x *TreeNode, sum *int) {
	if x != nil {
		InorderSum(x.Left, sum)
		*sum = *sum + x.Val
		InorderSum(x.Right, sum)
	}
}

func maxProduct(root *TreeNode) int {
	product := 0
	InorderSplit(root, root, &product)
	return product
}
