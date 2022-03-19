package _2156_find_substring_with_given_hash_value

import "testing"

func Test_subStrHash(t *testing.T) {
	tests := []struct {
		s         string
		power     int
		modulo    int
		k         int
		hashValue int
		want      string
	}{
		{"leetcode", 7, 20, 2, 0, "ee"},
		{"fbxzaad", 31, 100, 3, 32, "fbx"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := subStrHash(tt.s, tt.power, tt.modulo, tt.k, tt.hashValue); got != tt.want {
				t.Errorf("subStrHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
