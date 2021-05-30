package _43

import "testing"

func Test_isSumEqual(t *testing.T) {

	tests := []struct {
		firstWord  string
		secondWord string
		targetWord string
		want       bool
	}{
		//{"aacb", "cba", "cdb", true},
		{"aaa", "a", "aab", false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isSumEqual(tt.firstWord, tt.secondWord, tt.targetWord); got != tt.want {
				t.Errorf("isSumEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
