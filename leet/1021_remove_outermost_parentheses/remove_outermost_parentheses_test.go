package _1021_remove_outermost_parentheses

import "testing"

func Test_removeOuterParentheses(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{"1", "(()())(())", "()()()"},
		{"2", "(()())(())(()(()))", "()()()()(())"},
		{"3", "()()", ""},
		{"4", "", ""},
		{"5", "()", ""},
		{"5", "((()))", "(())"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOuterParentheses(tt.S); got != tt.want {
				t.Errorf("removeOuterParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
