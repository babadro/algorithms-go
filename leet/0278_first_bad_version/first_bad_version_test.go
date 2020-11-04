package _278_first_bad_version

import "testing"

const badVersion = 10

func Test_firstBadVersion(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"1", 7, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstBadVersion(tt.n); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
