package _1496_path_crossing

import "testing"

func Test_isPathCrossing(t *testing.T) {
	tests := []struct {
		path string
		want bool
	}{
		{"NNESWW", true},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			if got := isPathCrossing(tt.path); got != tt.want {
				t.Errorf("isPathCrossing() = %v, want %v", got, tt.want)
			}
		})
	}
}
