package _1249_minimum_remove_to_make_valid_parentheses

import "testing"

func Test_minRemoveToMakeValid(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"lee(t(c)o)de)", "lee(t(c)o)de"},
		{"a)b(c)d", "ab(c)d"},
		{"))((", ""},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := minRemoveToMakeValid3(tt.s); got != tt.want {
				t.Errorf("minRemoveToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
