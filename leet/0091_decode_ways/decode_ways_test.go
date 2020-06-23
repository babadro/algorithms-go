package _091_decode_ways

import "testing"

func TestNumDecodings(t *testing.T) {
	cases := []struct {
		s        string
		expected int
	}{
		{"12", 2},
		{"226", 3},
	}

	for i, c := range cases {
		if fact := numDecodings(c.s); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}
