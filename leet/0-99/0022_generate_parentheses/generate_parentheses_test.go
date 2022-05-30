package _022_generate_parentheses

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{3, []string{"((()))", "(()())", "()(())", "(())()", "()()()"}},
		{1, []string{"()"}},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			got := generateParenthesis2(tt.n)
			sort.Strings(got)
			sort.Strings(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
