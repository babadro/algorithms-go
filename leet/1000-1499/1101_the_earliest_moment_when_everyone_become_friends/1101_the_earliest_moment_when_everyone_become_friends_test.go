package _1101_the_earliest_moment_when_everyone_become_friends

import "testing"

func Test_earliestAcq(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		logs [][]int
		n    int
		want int
	}{
		//{
		//	[][]int{{20190101, 0, 1}, {20190104, 3, 4}, {20190107, 2, 3}, {20190211, 1, 5}, {20190224, 2, 4}, {20190301, 0, 3}, {20190312, 1, 2}, {20190322, 4, 5}},
		//	6,
		//	20190301,
		//},
		//{
		//	[][]int{{0, 2, 0}, {1, 0, 1}, {3, 0, 3}, {4, 1, 2}, {7, 3, 1}}, 4, 3,
		//},
		{
			[][]int{{9, 3, 0}, {0, 2, 1}, {8, 0, 1}, {1, 3, 2}, {2, 2, 0}, {3, 3, 1}}, 4, 2,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := earliestAcq(tt.logs, tt.n); got != tt.want {
				t.Errorf("earliestAcq() = %v, want %v", got, tt.want)
			}
		})
	}
}
