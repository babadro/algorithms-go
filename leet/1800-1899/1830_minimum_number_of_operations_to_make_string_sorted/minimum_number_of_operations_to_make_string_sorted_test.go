package _1830_minimum_number_of_operations_to_make_string_sorted

import "testing"

func Test_makeStringSorted(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"cba", 5},
		{"aabaa", 2},
		{"cdbea", 63},
		{"leetcodeleetcodeleetcode", 982157772},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := makeStringSorted2(tt.s); got != tt.want {
				t.Errorf("makeStringSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	tests := []struct {
		s    string
		i    int
		want string
	}{
		{"abc", 0, "cba"},
		{"abcdefg", 2, "abgfedc"},
		{"abcdefg", 6, "abcdefg"},
		{"abcdefg", 5, "abcdegf"},
		{"abcdefg", 4, "abcdgfe"},
		{"abcdefg", 3, "abcgfed"},
		{"abcdefg", 1, "agfedcb"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := reverse(tt.s, tt.i); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
