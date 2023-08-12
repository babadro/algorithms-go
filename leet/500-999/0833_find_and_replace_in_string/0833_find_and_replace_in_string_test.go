package _833_find_and_replace_in_string

import "testing"

func Test_findReplaceString(t *testing.T) {
	tests := []struct {
		s       string
		indices []int
		sources []string
		targets []string
		want    string
	}{
		{s: "abcd", indices: []int{0, 2}, sources: []string{"a", "cd"},
			targets: []string{"eee", "ffff"}, want: "eeebffff"},
		{s: "abcd", indices: []int{0, 2}, sources: []string{"ab", "ec"},
			targets: []string{"eee", "ffff"}, want: "eeecd"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := findReplaceString(tt.s, tt.indices, tt.sources, tt.targets); got != tt.want {
				t.Errorf("findReplaceString() = %v, want %v", got, tt.want)
			}
		})
	}
}
