package _219_contains_duplicate_2

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{
			name: "1",
			nums: []int{1, 2, 3, 1},
			k:    3,
			want: true,
		},
		{
			name: "2",
			nums: []int{1, 0, 1, 1},
			k:    1,
			want: true,
		},
		{
			name: "3",
			nums: []int{1, 2, 3, 1, 2, 3},
			k:    2,
			want: false,
		},
		{
			name: "4",
			nums: []int{},
			k:    0,
			want: false,
		},
		{
			name: "5",
			nums: []int{1},
			k:    0,
			want: false,
		},
		{
			name: "6",
			nums: []int{1, 1},
			k:    1,
			want: true,
		},
		{
			name: "7",
			nums: []int{1, 2, 1},
			k:    1,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.nums, tt.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
