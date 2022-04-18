package _3_string_interleaving

import (
	"testing"
)

func Test_findSIBruteForce(t *testing.T) {
	tests := []struct {
		m    string
		n    string
		p    string
		want bool
	}{
		{"abd", "cef", "abcdef", true},
		{"abd", "cef", "adcbef", false},
		{"abc", "def", "abdccf", false},
		{"abcdef", "mnop", "mnaobcdepf", true},
		{"aba", "aba", "abaaba", true},
		{"aba", "aba", "aba", false},
		{"aaaaaaaaaaaa", "bbbbbbbbbbbb", "abababababababababababab", true},
	}
	for _, tt := range tests {
		t.Run(tt.m+"_"+tt.n+"_"+tt.p, func(t *testing.T) {
			if got := findSIBottomUp(tt.m, tt.n, tt.p); got != tt.want {
				t.Errorf("findSI() = %v, want %v", got, tt.want)
			}
		})
	}
}

var m = "aaaaaaaaaaaa"
var n = "bbbbbbbbbbbbd"
var p = "ababababababababababababc"

// top down slower than bruteforce
func BenchmarkTopDown(b *testing.B) {
	var res bool
	for i := 0; i < b.N; i++ {
		res = findSITopDown(m, n, p)
	}
	_ = res
}
