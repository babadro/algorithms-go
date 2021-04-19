package _528_shuffle_string

import "testing"

func Test_restoreString(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		indices []int
		want    string
	}{
		{"1", "codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}, "leetcode"},
		{"2", "aiohn", []int{3, 1, 4, 2, 0}, "nihao"},
		{"3", "aaiougrt", []int{4, 0, 2, 6, 7, 3, 1, 5}, "arigatou"},
		{"4", "art", []int{1, 0, 2}, "rat"},
		{"4", "a", []int{0}, "a"},
		{"4", "пр", []int{1, 0}, "рп"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreString2(tt.s, tt.indices); got != tt.want {
				t.Errorf("restoreString() = %v, want %v", got, tt.want)
			}
		})
	}
}
