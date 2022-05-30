package _022_generate_parentheses

// passed. tptl. recursive
func generateParenthesis(n int) []string {
	var res []string
	rec(0, 0, make([]byte, n*2), &res)

	return res
}

func rec(i, balance int, cur []byte, res *[]string) {
	if i == len(cur) {
		*res = append(*res, string(cur))
		return
	}

	if balance > 0 {
		cur[i] = ')'
		rec(i+1, balance-1, cur, res)
	}

	if balance < len(cur)-i {
		cur[i] = '('
		rec(i+1, balance+1, cur, res)
	}
}

// passed. iterative. don't like it because of wasting resources for queue, perm structures etc
func generateParenthesis2(n int) (res []string) {
	q := []permutation{{}}
	var perm permutation
	for len(q) > 0 {
		perm, q = q[0], q[1:]
		if len(perm.val) == n*2 {
			res = append(res, string(perm.val))
		}

		if perm.balance > 0 {
			q = append(q, permutation{
				balance: perm.balance - 1,
				val:     perm.val + ")",
			})
		}

		if perm.balance < n*2-len(perm.val) {
			q = append(q, permutation{
				balance: perm.balance + 1,
				val:     perm.val + "(",
			})
		}
	}

	return
}

type permutation struct {
	balance int
	val     string
}
