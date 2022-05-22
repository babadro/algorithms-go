package _451_sort_characters_by_frequency

import "testing"

func Test_frequencySort(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"loveleetcode", "eeeeoollvtdc"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := frequencySort(tt.s); got != tt.want {
				t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
