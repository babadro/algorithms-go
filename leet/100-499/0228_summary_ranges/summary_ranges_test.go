package _228_summary_ranges

import (
	"reflect"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		wantRes []string
	}{
		{
			name:    "1",
			nums:    []int{0, 1, 2, 4, 5, 7},
			wantRes: []string{"0->2", "4->5", "7"},
		},
		{
			name:    "2",
			nums:    []int{0, 2, 3, 4, 6, 8, 9},
			wantRes: []string{"0", "2->4", "6", "8->9"},
		},
		{
			name:    "3",
			nums:    []int{-1},
			wantRes: []string{"-1"},
		},
		{
			name:    "4",
			nums:    []int{0},
			wantRes: []string{"0"},
		},
		{
			name:    "5",
			nums:    []int{-5, -3, -2, -1, 1, 2},
			wantRes: []string{"-5", "-3->-1", "1->2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := summaryRanges2(tt.nums); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("summaryRanges() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
