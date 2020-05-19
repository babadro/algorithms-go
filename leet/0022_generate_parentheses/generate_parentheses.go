package _022_generate_parentheses

// TODO 1
// брутфорс может быть таким - перебор всех возможных комбинаций - с проверкой на валидность.
func generateParenthesis(n int) []string {
	max := getMax(n)
	_ = max
	return nil
}

func getMax(n int) int {
	max := 1<<(n*2) - 1
	for i := 0; i < n; i++ {
		max &^= 1 << i
	}
	return max
}

func bitCount(n int) int {
	res := 0
	for n != 0 {
		if n&1 == 1 {
			res++
		}
		n = n >> 1
	}
	return res
}
