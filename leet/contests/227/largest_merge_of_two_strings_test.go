package _27

import "testing"

func Test_largestMerge(t *testing.T) {
	tests := []struct {
		word1 string
		word2 string
		want  string
	}{
		{"cabaa", "bcaaa", "cbcabaaaaa"},
		{"abcabc", "abdcaba", "abdcabcabcaba"},
		{"qqqqqqqqqeqeqqeeqqq", "qqqqqqqqeqqqeeqqeeqqqqqeq", "qqqqqqqqqqqqqqqqqeqqqeqeqqeeqqqeeqqeeqqqqqeq"},
	}
	for _, tt := range tests {
		t.Run(tt.word1+"_"+tt.word2, func(t *testing.T) {
			if got := largestMerge(tt.word1, tt.word2); got != tt.want {
				t.Errorf("largestMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

//"qqqqqqqqqeqeqqeeqqq"
//"qqqqqqqqeqqqeeqqeeqqqqqeq"
//"qqqqqqqqqqqqqqqqqeqqqeqeqqeeqqqeeqqeeqqqqqeq"
