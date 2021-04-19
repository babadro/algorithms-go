package _020_valid_parentheses

import "testing"

func TestValidate(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"{", false},
		{"{}", true},
		{"((())){}[]{{}}([])", true},
		{"(({)})", false},
		{"([)]", false},
		{"{[]}", true},
		{"((", false},
		{"[(({})}]", false},
		{"[([]])", false},
		{"(()(", false},
	}

	for i, v := range cases {
		if fact := isValid(v.input); fact != v.expected {
			t.Errorf("case#%d: %q, want %t, got %t", i+1, v.input, v.expected, fact)
		}
	}
}
