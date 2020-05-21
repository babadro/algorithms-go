package _022_generate_parentheses

import "testing"

func TestGenerateParentheses(t *testing.T) {
	//t.Log(getMax(1))
	//t.Log(bitCount(56))
	for _, str := range generateParenthesis(2) {
		t.Log(str)
	}
}
