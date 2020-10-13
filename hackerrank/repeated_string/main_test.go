package main

import "testing"

func Test_repeatedString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		n    int64
		want int64
	}{
		{
			"1", "aba", 11, 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedString(tt.s, tt.n); got != tt.want {
				t.Errorf("repeatedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
