package perm

import (
	"reflect"
	"testing"
)

func TestNextPermutationBytes(t *testing.T) {
	tests := []struct {
		arr     string
		wantArr string
		wantOk  bool
	}{
		{"1234", "1243", true},
		{"4321", "1234", false},
		{"1", "1", false},
		{"12", "21", true},
		{"158476531", "158513467", true},
	}
	for _, tt := range tests {
		t.Run(tt.arr, func(t *testing.T) {
			arr := []byte(tt.arr)
			gotOk := NextPermutationBytes(arr)
			if gotOk != tt.wantOk {
				t.Errorf("NextPermutationBytes() = %v, want %v", gotOk, tt.wantOk)
			} else if !reflect.DeepEqual(arr, []byte(tt.wantArr)) {
				t.Errorf("want %s, got %s", tt.wantArr, string(arr))
			}
		})
	}
}
