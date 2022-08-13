package _091_decode_ways

import "testing"

func TestNumDecodings(t *testing.T) {
	cases := []struct {
		s        string
		expected int
	}{
		{"12", 2},
		{"226", 3},
		{"351226", 5},
		{"0", 0},
		{"10", 1},
		{"1010", 1},
		{"10210", 1},
		{"01", 0},
		{"20101", 1},
		{"230", 0},
		{"17", 2},
		{"1090", 0},
		{"102", 1},
		{"22634212112312111234212311211231131112123112121232", 297555930},
		{"111111111111111111111111111111111111111111111", 1836311903},
	}

	for i, c := range cases {
		if fact := numDecodings2(c.s); fact != c.expected {
			t.Errorf("case#%d, want %d, got %d", i+1, c.expected, fact)
		}
	}
}
