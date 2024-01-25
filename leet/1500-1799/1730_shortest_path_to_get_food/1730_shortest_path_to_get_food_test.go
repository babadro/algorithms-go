package _1730_shortest_path_to_get_food

import "testing"

func Test_getFood(t *testing.T) {
	tests := []struct {
		grid [][]byte
		want int
	}{
		{
			grid: [][]byte{
				[]byte("XXXXXX"),
				[]byte("X*000X"),
				[]byte("X00#0X"),
				[]byte("XXXXXX"),
			},
			want: 3,
		},
		{
			grid: [][]byte{
				[]byte("XXXXX"),
				[]byte("X*X0X"),
				[]byte("X0X#X"),
				[]byte("XXXXX"),
			},
			want: -1,
		},
		{
			grid: [][]byte{
				[]byte("XXXXXXXX"),
				[]byte("X*0X0#0X"),
				[]byte("X00X00XX"),
				[]byte("X0000#0X"),
				[]byte("XXXXXXXX"),
			},
			want: 6,
		},
		{
			grid: [][]byte{
				{'#', '0', '0', '0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0', '0', '0', '0'},
				{'0', '0', '#', '#', '0', '0', '0', '0'},
				{'0', '0', '#', '0', 'X', '0', '0', '0'},
				{'0', '0', '0', '0', '0', '0', '#', '0'},
				{'0', '0', '0', '#', '*', '0', '0', '0'},
				{'0', '0', '0', '0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			grid: [][]byte{},
		},
		{
			grid: [][]byte{
				[]byte("000"),
				[]byte("0##"),
				[]byte("00X"),
				[]byte("000"),
				[]byte("#0*"),
			},
			want: 2,
		},
		{
			grid: tle, want: -1,
		},
		{
			grid: tle2, want: 1,
		},
		{grid: tle3, want: 198},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := getFood(tt.grid); got != tt.want {
				t.Errorf("getFood() = %v, want %v", got, tt.want)
			}
		})
	}
}
