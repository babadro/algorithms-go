package _1830_minimum_number_of_operations_to_make_string_sorted

import "testing"

func Test_makeStringSorted(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"cba", 5},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := makeStringSorted(tt.s); got != tt.want {
				t.Errorf("makeStringSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
