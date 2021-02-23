package _1417_reformat_the_string

import (
	"sort"
	"testing"
)

func Test_reformat(t *testing.T) {
	tests := []struct {
		s     string
		empty bool
	}{
		{"a0b1c2", false},
		//{"leetcode", true},
		//{"1229857369", true},
		//{"covid2019", false},
		//{"ab123", false},
		//{"a", false},
		//{"a1", false},
		//{"1", false},
		//{"1a", false},
		//{"1a1", false},
		//{"11a", false},
		//{"a11", false},
		//{"aa1", false},
		//{"a1a", false},
		//{"1aa", false},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := reformat(tt.s)
			if !tt.empty {
				b := []byte(got)
				sort.Slice(b, func(i, j int) bool {
					return b[i] < b[j]
				})
				input := []byte(tt.s)
				sort.Slice(input, func(i, j int) bool {
					return input[i] < input[j]
				})

				pass := true
				for i := 1; i < len(got); i++ {
					if digit(got[i]) && digit(got[i-1]) {
						pass = false
						break
					}
				}

				if string(b) != string(input) || !pass {
					t.Errorf("reformat() = %v", reformat(tt.s))
				}

			} else if got != "" {
				t.Errorf("want empty string, but got %s", got)
			}
		})
	}
}
