package _5_alien_dictionary

import "testing"

func Test_findOrder(t *testing.T) {
	tests := []struct {
		words []string
		want  string
	}{
		{[]string{"ba", "bc", "ac", "cab"}, "bac"},
		{[]string{"cab", "aaa", "aab"}, "cab"},
		{[]string{"ywx", "wz", "xww", "xz", "zyy", "zwz"}, "ywxz"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findOrder(tt.words); got != tt.want {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
