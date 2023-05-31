package _2101_detonate_the_maximum_bombs

import "testing"

func Test_maximumDetonation(t *testing.T) {
	tests := []struct {
		bombs [][]int
		want  int
	}{
		{[][]int{{2, 1, 3}, {6, 1, 4}}, 2},
		{[][]int{{1, 1, 5}, {10, 10, 5}}, 1},
		{
			[][]int{{855, 82, 158}, {17, 719, 430}, {90, 756, 164}, {376, 17, 340}, {691, 636, 152}, {565, 776, 5}, {464, 154, 271}, {53, 361, 162}, {278, 609, 82}, {202, 927, 219}, {542, 865, 377}, {330, 402, 270}, {720, 199, 10}, {986, 697, 443}, {471, 296, 69}, {393, 81, 404}, {127, 405, 177}},
			9,
		},
		{
			[][]int{{1, 1, 5}, {10, 10, 5}}, 1,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maximumDetonation(tt.bombs); got != tt.want {
				t.Errorf("maximumDetonation() = %v, want %v", got, tt.want)
			}
		})
	}
}
