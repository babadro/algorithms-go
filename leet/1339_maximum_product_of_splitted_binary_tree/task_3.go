package _1339_maximum_product_of_splitted_binary_tree

import "github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"

// TODO Похоже, дело было в вычислении остатков. Надо понять почему.
// TODO - добавил новый большой файл с деревом - у меня timeLimitExceed - надо оптимизировать функцию, чтобы не считал
// каждый раз сумму дерева
func calcProduct(root1, root2 *btree.Node) int {
	var sum1 int
	f := func(node *btree.Node) { sum1 += node.Val }
	btree.InOrderFunc(root1, f)

	var sum2 int
	f = func(node *btree.Node) { sum2 += node.Val }
	btree.InOrderFunc(root2, f)

	return sum1 * sum2

	// Похоже, вот тут была ошибка. Надо понять, почему...
	i := 1
	m := 1000000007
	i = (i * sum1) % m
	i = (i * sum2) % m
	return i
}

func maxProduct2(root *btree.Node) int {
	var product int
	f := func(x *btree.Node) {
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
	btree.InOrderFunc(root, f)
	return product % (int(1e9) + 7)
}
